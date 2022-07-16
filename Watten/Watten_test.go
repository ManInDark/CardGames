package Watten

import (
	"CardGames/CardDeck"
	"fmt"
	"testing"
)

func TestKritische(t *testing.T) {
	kritische := kritische()
	if kritische[0] != CardDeck.CreateCard(CardDeck.HERZ, CardDeck.KÖNIG) {
		t.Fatal("kritische function not working properly")
	} else if kritische[1] != CardDeck.CreateCard(CardDeck.SCHELLE, CardDeck.SEVEN) {
		t.Fatal("kritische function not working properly")
	} else if kritische[2] != CardDeck.CreateCard(CardDeck.EICHEL, CardDeck.SEVEN) {
		t.Fatal("kritische function not working properly")
	}
}

func TestCreateWatten(t *testing.T) {
	watten := CreateWatten(true, true)
	if len(watten.Players) != 4 {
		t.Fatal("CreateWatten function not working properly")
	} else if watten.Turn != 0 {
		t.Fatal("CreateWatten function not working properly")
	} else if watten.Large {
		t.Fatal("CreateWatten function not working properly")
	}
}

func TestFindWinner(t *testing.T) {
	if findWinner([]CardDeck.Card{ // Kritter
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS),
		CardDeck.CreateCard(CardDeck.EICHEL, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.KÖNIG)},
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.NINE)) != 3 {
		t.Fatal("findWinner not working properly")
	}
	if findWinner([]CardDeck.Card{ // Haube
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.NINE)},
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.NINE)) != 3 {
		t.Fatal("findWinner not working properly")
	}
	if findWinner([]CardDeck.Card{ // Schlag
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.NINE)},
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.NINE)) != 3 {
		t.Fatal("findWinner not working properly")
	}
	if findWinner([]CardDeck.Card{ // Farbe
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.TEN)},
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.NINE)) != 3 {
		t.Fatal("findWinner not working properly")
	}
	winner := findWinner([]CardDeck.Card{ // Erste Farbe
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.TEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.BUBE),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.SEVEN),
		CardDeck.CreateCard(CardDeck.HERZ, CardDeck.ASS)},
		CardDeck.CreateCard(CardDeck.BLATT, CardDeck.NINE))
	if winner != 3 {
		fmt.Println(winner)
		t.Fatal("findWinner not working properly")
	}
}
func TestWhole(t *testing.T) {
	watten := CreateWatten(false, true)
	go func() {
		for {
			msg := <-watten.Players[0].Stdout
			fmt.Println("Player 0:", msg)
		}
	}()
	go func() {
		watten.Players[0].Stdin <- "apfelsaft"
		watten.Players[0].Stdin <- "eichel"
		watten.Players[0].Stdin <- "0"
		watten.Players[0].Stdin <- "0"
		watten.Players[0].Stdin <- "0"
		watten.Players[0].Stdin <- "0"
		watten.Players[0].Stdin <- "0"
	}()
	go func() {
		for {
			msg := <-watten.Players[1].Stdout
			fmt.Println("Player 1:", msg)
		}
	}()
	go func() {
		watten.Players[1].Stdin <- "braten"
		watten.Players[1].Stdin <- "acht"
		watten.Players[1].Stdin <- "0"
		watten.Players[1].Stdin <- "0"
		watten.Players[1].Stdin <- "0"
		watten.Players[1].Stdin <- "0"
		watten.Players[1].Stdin <- "0"
	}()
	watten.RunRound()
}
