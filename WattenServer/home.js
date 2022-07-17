socket = new WebSocket("ws://" + location.host + "/socket")
socket.onerror = error => { console.log("Socket Error: ", error) }
socket.onmessage = message => {
    console.log(message.data)
}
function send(message) { socket.send(message) }
function list() { send("list") }