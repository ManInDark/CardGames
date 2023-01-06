package Schafkopfen

import (
	"CardGames/CardDeck"
	"fmt"
	"testing"
)

func TestCreateSchafkopfen(t *testing.T) {
	schafkopfen := CreateSchafkopfen(true)
	if len(schafkopfen.Players) != 4 {
		t.Fatal("CreateWatten function not working properly")
	} else if schafkopfen.Turn != 0 {
		t.Fatal("CreateWatten function not working properly")
	}
}

func TestFindWinner(t *testing.T) {
	if findWinner([]CardDeck.Card{ // Geier
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS),
		CardDeck.CreateCard(CardDeck.EICHEL, CardDeck.OBER),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.KÖNIG)},
		Spielweise{[]*CardDeck.Player{}, 0, GEIER}) != 2 {
		t.Fatal("findWinner not working properly")
	}
	if findWinner([]CardDeck.Card{ // Wenz
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.UNTER),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.OBER)},
		Spielweise{[]*CardDeck.Player{}, 0, WENZ}) != 2 {
		t.Fatal("findWinner not working properly")
	}
	if findWinner([]CardDeck.Card{ // Solo
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.EICHEL, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.NINE)},
		Spielweise{[]*CardDeck.Player{}, CardDeck.EICHEL, SOLO}) != 1 {
		t.Fatal("findWinner not working properly")
	}
	if findWinner([]CardDeck.Card{ // Ramsch / Erste Farbe
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.KÖNIG),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.EIGHT),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.TEN)},
		Spielweise{[]*CardDeck.Player{}, 0, RAMSCH}) != 3 {
		t.Fatal("findWinner not working properly")
	}
	if findWinner([]CardDeck.Card{ // Herz Stich
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.EIGHT),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.SEVEN)},
		Spielweise{[]*CardDeck.Player{}, 0, RAMSCH}) != 3 {
		t.Fatal("findWinner not working properly")
	}
}
func TestWhole(t *testing.T) {
	watten := CreateSchafkopfen(true)
	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				msg := <-watten.Players[i].Stdout
				fmt.Println(fmt.Sprintf("Player %d:", i), msg)
			}
		}(i)
		go func(i int) {
			watten.Players[i].Stdin <- "apfelsaft"
			watten.Players[i].Stdin <- "Ramsch"
			for n := 0; n < 8; n++ {
				watten.Players[i].Stdin <- "0"
			}
		}(i)
	}
	watten.RunRound()
}
