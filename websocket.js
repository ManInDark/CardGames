socket0 = new WebSocket("ws://server-debian.local:2000")
socket0.onerror = error => { console.log("Socket Error: ", error) }
socket0.onmessage = message => {
    console.log("Player 0: " + message.data)
}
socket1 = new WebSocket("ws://server-debian.local:2000")
socket1.onerror = error => { console.log("Socket Error: ", error) }
socket1.onmessage = message => {
    console.log("Player 1: " + message.data)
}

socket1.send("acht")
socket0.send("herz")

socket1.send("0")
socket1.send("0")
socket1.send("0")
socket1.send("0")
socket1.send("0")

socket0.send("0")
socket0.send("0")
socket0.send("0")
socket0.send("0")
socket0.send("0")