socket = new WebSocket("ws://" + location.host + "/socket")
socket.onerror = error => { console.log("Socket Error: ", error) }
socket.onmessage = message => {
    console.log(message.data)
}
function send(message) { socket.send(message) }
function list() { send("list") }

function checkKritter(value, color) {
    value = value.toUpperCase();
    color = color.toUpperCase();
    if (color === "HERZ" && value === "KÖNIG")
        return "GoldenRod";
    else if (color === "SCHELLE" && value === "SEVEN")
        return "silver";
    else if (color === "EICHEL" && value === "SEVEN")
        return "sienna";
    else
        return "none";
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
    }
}

class GameCard extends HTMLElement {
    constructor() {
        super();
    }

    connectedCallback() {
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

customElements.define("game-card", GameCard);