package main

import (
	"CardGames/CardDeck"
	"CardGames/Schafkopfen"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var games = []*Schafkopfen.Schafkopfen{}

func websocketConnector(w *http.ResponseWriter, r *http.Request, game *Schafkopfen.Schafkopfen) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	socket, err := upgrader.Upgrade(*w, r, nil)

	if err != nil {
		return
	}

	player := CardDeck.CreatePlayer("Player") // TODO player namen irgendwie festlegen
	game.Players = append(game.Players, &player)
	fmt.Println("Connection Received, now " + strconv.Itoa(len(game.Players)) + " Players")

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

	if len(game.Players) >= 4 {
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
	selected_game := games[len(games)-1]
	if len(selected_game.Players) < 4 {
		websocketConnector(&w, r, selected_game)
	} else {
		newgame := Schafkopfen.CreateSchafkopfen(false)
		games = append(games, &newgame)
		websocketConnector(&w, r, &newgame)
	}
}

func main() {
	// Initialize first game
	newgame := Schafkopfen.CreateSchafkopfen(false)
	games = append(games, &newgame)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home.html", http.StatusMovedPermanently)
	})

	http.HandleFunc("/home.html", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "home.html") })
	http.HandleFunc("/home.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "home.js") })

	http.HandleFunc("/blatt.svg", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "svg/blatt.svg") })
	http.HandleFunc("/eichel.svg", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "svg/eichel.svg") })
	http.HandleFunc("/herz.svg", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "svg/herz.svg") })
	http.HandleFunc("/schelle.svg", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "svg/schelle.svg") })

	http.HandleFunc("/socket", websocketHandler)
	http.ListenAndServe(":2000", nil)
}
