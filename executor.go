package main

import (
	"CardGames/CardDeck"
	"fmt"
)

func main() {
	deck := CardDeck.CreateDeck(CardDeck.SIX)
	deck.Shuffle()
	deck.Lift()
	fmt.Println(deck.Peek(-1))
	fmt.Println(deck.Peek(-1) == deck.Take(-1))
	card := deck.Take(-1)

	fmt.Println(CardDeck.ASS > CardDeck.KÖNIG)
	fmt.Println(card.GetColor(), card.GetValue())

	player := CardDeck.CreatePlayer("somename")
	player.AddCard(deck.Take(0))
	fmt.Println(player)
	fmt.Println(CardDeck.ListToString(player.ListCards()))
	fmt.Println(player.GetCard(0))
	fmt.Println(CardDeck.ListToString(player.ListCards()))
}
