package Watten

import (
	"CardGames/CardDeck"
	"fmt"
	"strconv"
	"strings"
)

const (
	DEFAULT_SIZE int = 2
)

func kritische() []CardDeck.Card {
	return []CardDeck.Card{
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.KÖNIG),
		CardDeck.CreateCard(CardDeck.SCHELLE, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.EICHEL, CardDeck.SEVEN),
	}
}

type Watten struct {
	Turn    int
	Deck    CardDeck.Deck
	Players []*CardDeck.Player
	Large   bool
}

func (watten *Watten) writeOutputAll(message string) {
	fmt.Println(message)
	for _, player := range watten.Players {
		player.Stdout <- message
	}
}

func CreateWatten(large bool, with_players bool) Watten {
	watten := Watten{0, CardDeck.CreateDeck(CardDeck.SIX), []*CardDeck.Player{}, false}
	playercount := DEFAULT_SIZE
	if with_players {
		if large {
			playercount *= 2
		}
		for n := 0; n < playercount; n++ {
			player := (CardDeck.CreatePlayer(string(rune(n))))
			watten.Players = append(watten.Players, &player)
		}
	}
	return watten
}

type IntCard struct {
	n    int
	card CardDeck.Card
}

func findWinner(cards []CardDeck.Card, haube CardDeck.Card) int {
	// Kritter
	for _, kritter := range kritische() {
		for n, card := range cards {
			if card == kritter {
				return n
			}
		}
	}
	// Haube
	for n, card := range cards {
		if card == haube {
			return n
		}
	}
	// Schlag
	for n, card := range cards {
		if card.GetValue() == haube.GetValue() {
			return n
		}
	}
	// Farbe
	var temp_card IntCard
	for n, card := range cards {
		if card.GetColor() == haube.GetColor() && card.GetValue() > temp_card.card.GetValue() {
			temp_card = IntCard{n, card}
		}
	}
	if temp_card.card.GetValue().String() != "" {
		return temp_card.n
	}
	// Erste Farbe
	temp_card = IntCard{0, cards[0]}
	for n, card := range cards {
		if card.GetColor() == temp_card.card.GetColor() && card.GetValue() > temp_card.card.GetValue() {
			temp_card = IntCard{n, card}
		}
	}
	return temp_card.n
}

func requestSchlag(player *CardDeck.Player, sw_possible bool) (CardDeck.Value, bool) {
	if sw_possible {
		player.Stdout <- "Gewünschter Schlag ist:"
	} else {
		player.Stdout <- "Gewünschter Schlag ist: (Kein Schlagwechsel möglich)"
	}
	for {
		response := strings.Trim(<-player.Stdin, " \t\n\r")
		if strings.EqualFold(response, "Schlagwechsel") {
			return CardDeck.SIX, true
		}
		for value := CardDeck.SIX; value <= CardDeck.ASS; value++ {
			if strings.EqualFold(value.String(), response) {
				return value, false
			}
		}
		player.Stdout <- "Auswahl nicht erkannt"
	}
}

func requestFarbe(player *CardDeck.Player) CardDeck.Color {
	player.Stdout <- "Gewünschte Farbe ist:"
	for {
		response := strings.Trim(<-player.Stdin, " \t\n\r")
		for color := CardDeck.EICHEL; color <= CardDeck.BLATT; color++ {
			if strings.EqualFold(color.String(), response) {
				return color
			}
		}
		player.Stdout <- "Auswahl nicht erkannt"
	}
}

func acceptSchlagwechsel(player *CardDeck.Player) bool {
	player.Stdout <- "Schlagwechsel annehmen?"
	for {
		response := strings.Trim(<-player.Stdin, " \t\n\r")
		if strings.EqualFold(response, "Ja") {
			return true
		} else if strings.EqualFold(response, "Nein") {
			return false
		}
		player.Stdout <- "Auswahl nicht erkannt"
	}
}

func (watten *Watten) RunRound() {
	watten.Deck.Shuffle()
	watten.Deck.Lift()
	watten.writeOutputAll("Die Runde hat begonnen")

	// Abgehobene Karte ansehen
	abgehobene := watten.Deck.Peek(-1)
	watten.Players[(watten.Turn+3)%len(watten.Players)].Stdout <- "Abgehobene Karte: " + abgehobene.String()
	for _, kritischer := range kritische() {
		if abgehobene == kritischer {
			watten.Players[(watten.Turn+3)%len(watten.Players)].Stdout <- "Karte genommen"
			watten.Players[(watten.Turn+3)%len(watten.Players)].AddCard(watten.Deck.Take(-1))
		}
	}

	// Restliche Karten austeilen
	for _, player := range watten.Players {
		remcards := 3 - len(player.ListCards())
		for n := 0; n < remcards; n++ {
			player.AddCard(watten.Deck.Take(0))
		}
	}
	for _, player := range watten.Players {
		for n := 0; n < 2; n++ {
			player.AddCard(watten.Deck.Take(0))
		}
	}

	fp := watten.Players[(watten.Turn+1)%len(watten.Players)]
	sp := watten.Players[watten.Turn]
	// Schlag und Farbe ansagen
	value, schlagwechsel := requestSchlag(fp, true)
	if schlagwechsel {
		if acceptSchlagwechsel(sp) {
			fp, sp = sp, fp
		}
		for schlagwechsel { // Damit der andere nicht versucht das zurückzutauschen
			value, schlagwechsel = requestSchlag(fp, false)
		}
	}
	watten.writeOutputAll("Gewählter Schlag ist: " + value.String())
	color := requestFarbe(sp)
	watten.writeOutputAll("Gewählte Farbe ist: " + color.String())

	// Spielrunden
	punktestand := []int{0, 0}
	beginner := watten.Turn + 1
	haube := CardDeck.CreateCard(color, value)

	for n := 0; n < 5; n++ {
		gelegte_karten := []CardDeck.Card{}

		// jeweils die gelegten Karten auswählen
		for i := 0; i < len(watten.Players); i++ {
			player := watten.Players[(watten.Turn+i+beginner)%len(watten.Players)]
			player.Stdout <- "Zu legende Karte:"
			for {
				number, err := strconv.Atoi(strings.Trim(<-player.Stdin, " \t\n\r"))
				if number >= len(player.ListCards()) {
					player.Stdout <- "Invalide Karte"
					continue
				}
				if err == nil {
					gelegte_karten = append(gelegte_karten, player.GetCard(int8(number)))
					watten.writeOutputAll(gelegte_karten[i].String())
					break
				}
			}
		}
		winner := findWinner(gelegte_karten, haube)
		watten.writeOutputAll("Gewonnen hat: " + gelegte_karten[winner].String())
		punktestand[(beginner+winner)%len(watten.Players)] += 1
		beginner = beginner + winner
		watten.writeOutputAll("Zwischenstand: " + strconv.Itoa(punktestand[0]) + ":" + strconv.Itoa(punktestand[1]))
	}
	watten.writeOutputAll("Endgültiger Punktestand: " + strconv.Itoa(punktestand[0]) + ":" + strconv.Itoa(punktestand[1]))
	watten.Turn += 1
	watten.Turn %= len(watten.Players)
}
