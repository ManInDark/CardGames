var removal_timeout = 1000;
var state = "initializing"; // (-> selecting:value -> selecting:color) -> choosing -> playing
var haube = "";
socket = new WebSocket("ws" + location.protocol.replace("http", "") + "//" + location.host + "/socket");
socket.onerror = error => { console.log("Socket Error: ", error); }
socket.onmessage = message => {
    if (message.data.startsWith("[[")) { // liest list antwort ein
        message.data.replace("[[", "").replace("]]", "").split("] [").forEach(card => {
            split = card.split(" ");
            document.getElementById("holder").append(createCard(split[1], split[0]));
        });
    }
    else if (message.data.startsWith("[")) {
        split = message.data.replace("[", "").replace("]", "").split(" ");
        document.getElementById("cards").append(createCard(split[1].trim(), split[0].trim()));
    }
    else if (state === "initializing" && message.data === "Die Runde hat begonnen") { state = "choosing"; list(); haube = ""; }
    else if (state === "choosing" && message.data.startsWith("Gewünschte")) {
        state = "selecting:" + (message.data.split(" ")[1] === "Schlag" ? "value" : "color");
        if (message.data.split(" ")[1] === "Schlag" && !message.data.includes("Kein")) { document.getElementById("cards").append(createCard("Schlagwechsel", "Schlagwechsel")) }
        createLog("Wähle " + message.data.split(" ")[1] + " aus!");
    } else if (state === "choosing" && message.data === "Schlagwechsel annehmen?") { 
        document.getElementById("cards").append(createCard("Ja", "Ja"));
        document.getElementById("cards").append(createCard("Nein", "Nein"));
        createLog(message.data);
        state = "selecting:value";
    }
    else if (state === "choosing" && message.data === "Zu legende Karte:") { state = "playing"; createLog("Du bist dran!") }
    else if (message.data.startsWith("Gewählte")) {
        createLog(message.data);
        haube += message.data.split(": ")[1] + " ";
        if (haube.split(" ").length > 2) { // Wenn beide gewählt wurden
            split = haube.split(" ");
            document.getElementById("haube").append(createCard(split[0], split[1]));
        }
    }
    else if (message.data.startsWith("Gewonnen hat:")) {
        for (i = 0; i < document.getElementById("cards").children.length; i++)
            setTimeout((card) => { card.remove() }, removal_timeout, document.getElementById("cards").children[i])
    } else {
        createLog(message.data);
    }
    if (message.data.startsWith("Endgültiger")) { setTimeout(() => { document.getElementById("haube").children[0].remove() }, removal_timeout); }
    console.log(message.data);
}
function send(message) { socket.send(message) }
function list() { send("list") }

function checkKritter(value, color) {
    if (color === "Herz" && value === "König")
        return "GoldenRod";
    else if (color === "Schelle" && value === "Sieben")
        return "silver";
    else if (color === "Eichel" && value === "Sieben")
        return "sienna";
    else
        return "white";
}

function translateValue(value) {
    switch (value) {
        case "Sechs":
            return "6";
        case "Sieben":
            return "7";
        case "Acht":
            return "8";
        case "Neun":
            return "9";
        case "Zehn":
            return "10";
        case "Unter":
            return "U";
        case "Ober":
            return "O";
        case "König":
            return "K";
        case "Ass":
            return "A";
        case "Ja":
            return "✓";
        case "Nein":
            return "✗";
        case "Schlagwechsel":
            return "🗘";
    }
}

class GameCard extends HTMLElement {
    constructor() {
        super();
    }

    connectedCallback() {
        this.addEventListener("click", clickHandler);
        const svgElement = document.createElementNS("http://www.w3.org/2000/svg", "svg");
        svgElement.setAttribute("version", "1.0");
        svgElement.setAttribute("width", "100px");
        svgElement.setAttribute("height", "200px");
        this.appendChild(svgElement);

        const border = document.createElementNS("http://www.w3.org/2000/svg", "rect");
        border.setAttribute("x", "1px");
        border.setAttribute("y", "1px");
        border.setAttribute("width", "98px");
        border.setAttribute("height", "198px");
        border.setAttribute("stroke", "black");
        border.setAttribute("stroke-width", "2px");
        border.setAttribute("fill", checkKritter(this.getAttribute("value"), this.getAttribute("color")));
        border.setAttribute("rx", "10px");
        border.setAttribute("ry", "10px");
        svgElement.appendChild(border);

        const line = document.createElementNS("http://www.w3.org/2000/svg", "line");
        line.setAttribute("x1", "1");
        line.setAttribute("y1", "100");
        line.setAttribute("x2", "99");
        line.setAttribute("y2", "100");
        line.setAttribute("stroke", "black");
        line.setAttribute("stroke-width", "2px");
        svgElement.appendChild(line);

        const value = document.createElementNS("http://www.w3.org/2000/svg", "text");
        value.setAttribute("x", "50");
        value.setAttribute("y", "175");
        value.setAttribute("text-anchor", "middle");
        value.setAttribute("fill", "black");
        value.innerHTML = translateValue(this.getAttribute("value"));
        value.style.fontFamily = "Arial";
        value.style.fontSize = "4em";
        value.style.textAlign = "center";

        svgElement.appendChild(value);

        if (!["Ja", "Nein", "Schlagwechsel"].includes(this.getAttribute("color"))) {
            const color = document.createElementNS("http://www.w3.org/2000/svg", "image");
            color.setAttribute("x", "10");
            color.setAttribute("y", "10");
            color.setAttribute("width", "80");
            color.setAttribute("height", "80");
            // color.setAttribute("preserveAspectRatio", "meet");
            color.setAttribute("href", this.getAttribute("color").toLowerCase() + ".svg");
            svgElement.appendChild(color);
        }
    }
}

customElements.define("game-card", GameCard);

function createCard(value, color) {
    const newCard = document.createElement("game-card");
    newCard.setAttribute("value", value.trim());
    newCard.setAttribute("color", color.trim());
    return newCard;
}

function clickHandler() {
    if (state.startsWith("selecting")) { removeLastLog(); send(this.getAttribute(state.split(":")[1])); Array.from(document.getElementById("cards").children).forEach(element => element.remove()); state = "choosing"; }
    if (state === "playing") {
        for (i = 0; i < this.parentElement.childElementCount; i++) {
            if (this.parentElement.children[i] === this) {
                removeLastLog(); // Removes the log "Du bist dran"
                send(i);
                state = "choosing";
                this.remove();
                break;
            }
        }
    }
}

/**
 * creates a log element and appends it to the log 
 * 
 * @param {String} message 
 */
function createLog(message) {
    let logelement = document.createElement("p");
    logelement.innerText = message;
    document.getElementById("log").append(logelement);
}

/**
 * removes the last log message
 */
function removeLastLog() {
    document.getElementById("log").children[document.getElementById("log").childElementCount - 1].remove();
}