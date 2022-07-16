package main

import (
	"CardGames/CardDeck"
	"CardGames/Watten"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var games = []*Watten.Watten{}

func websocketConnector(w *http.ResponseWriter, r *http.Request, game *Watten.Watten) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	socket, err := upgrader.Upgrade(*w, r, nil)

	if err != nil {
		return
	}

	player := CardDeck.CreatePlayer("Player") // TODO player namen irgendwie festlegen
	game.Players = append(game.Players, &player)

	go func(socket *websocket.Conn, player *CardDeck.Player) {
		for {
			_, bytes, err := socket.ReadMessage()
			if err != nil {
				fmt.Println("Player disconnected because of error")
				return
			}
			received_message := strings.Trim(string(bytes), " \t\n\r")
			switch received_message {
			case "list":
				socket.WriteMessage(1, []byte(fmt.Sprintln(player.ListCards())))
			default:
				player.Stdin <- string(bytes)
			}
		}
	}(socket, &player)

	if len(game.Players)%2 == 0 && len(game.Players) > 2 == game.Large {
		go game.RunRound()
	}

	for { // write STDOUT to socket
		socket.WriteMessage(1, []byte(<-player.Stdout))
	}
}

/*
	Assigns new connections to games and creates new ones if necessary
	So just sorts where the connections should go, doesn't actually connect anything
*/
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connection Received")
	selected_game := games[len(games)-1]
	if selected_game.Large {
		if len(selected_game.Players) < 4 {
			websocketConnector(&w, r, selected_game)
		}
	} else {
		if len(selected_game.Players) < 2 {
			websocketConnector(&w, r, selected_game)
		}
	}
	newgame := Watten.CreateWatten(false, false)
	games = append(games, &newgame)
	websocketConnector(&w, r, &newgame)
}

func main() {
	// Initialize first game
	newgame := Watten.CreateWatten(false, false)
	games = append(games, &newgame)

	http.HandleFunc("/", websocketHandler)
	http.ListenAndServe(":2000", nil)
}
