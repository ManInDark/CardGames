package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Color int8
type Value int8

const (
	EICHEL Color = iota
	SCHELLE
	HERZ
	BLATT
	SIX   Value = 6
	SEVEN Value = 7
	EIGHT Value = 8
	NINE  Value = 9
	TEN   Value = 10
	BUBE  Value = 11
	DAME  Value = 12
	KÖNIG Value = 13
	ASS   Value = 14
)

func (color Color) String() string {
	switch color {
	case EICHEL:
		return "Eichel"
	case SCHELLE:
		return "Schelle"
	case HERZ:
		return "Herz"
	case BLATT:
		return "Blatt"
	}
	return "" // unreacheable
}

func (value Value) String() string {
	switch value {
	case SIX:
		return "Sechs"
	case SEVEN:
		return "Sieben"
	case EIGHT:
		return "Acht"
	case NINE:
		return "Neun"
	case TEN:
		return "Zehn"
	case BUBE:
		return "Bube"
	case DAME:
		return "Dame"
	case KÖNIG:
		return "König"
	case ASS:
		return "Ass"
	}
	return ""
}

type Card struct {
	color Color
	value Value
}

func (card Card) getColor() Color {
	return card.color
}

func (card Card) getValue() Value {
	return card.value
}

func (card Card) String() string {
	return "[" + card.color.String() + " " + card.value.String() + "]"
}

func listToString(cards []Card) string {
	str := "["
	for index, card := range cards {
		if index > 0 {
			str += ", "
		}
		str += card.String()
	}
	return str + "]"
}

type Player struct {
	cards  []Card
	name   string
	stdout chan string
	stdin  chan string
}

func (player Player) String() string {
	return player.name + ": " + listToString(player.cards)
}

func (player *Player) addCard(card Card) {
	player.cards = append(player.cards, card)
}

func (player *Player) getCard(index int8) Card {
	card := player.cards[index]
	shifted := player.cards[index+1:]
	player.cards = append(player.cards[:index], shifted...)
	return card
}

type Deck struct {
	cards []Card
}

func createDeck() Deck {
	deck := Deck{[]Card{}}
	for color := EICHEL; color <= BLATT; color++ {
		for value := SIX; value <= ASS; value++ {
			deck.cards = append(deck.cards, Card{color, value})
		}
	}
	return deck
}

func (deck Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })
}

func (deck *Deck) lift() {
	index := rand.Intn(len(deck.cards) - 1)
	lifted := deck.cards[:index]
	deck.cards = append(deck.cards[index:len(deck.cards)], lifted...)
}

func (deck Deck) peek(index int) Card {
	if index == -1 {
		index = len(deck.cards) - 1
	}
	return deck.cards[index]
}

func (deck *Deck) take(index int) Card {
	if index == -1 {
		index = len(deck.cards) - 1
	}
	card := deck.cards[index]
	shifted := deck.cards[index+1:]
	deck.cards = append(deck.cards[:index], shifted...)
	return card
}

func (deck Deck) String() string {
	return listToString(deck.cards)
}

func main() {
	fmt.Println(ASS > KÖNIG)
	card := Card{EICHEL, ASS}
	fmt.Println(card.getColor(), card.getValue())

	player := Player{[]Card{card}, "somename", make(chan string), make(chan string)}
	player.addCard(Card{HERZ, ASS})
	fmt.Println(listToString(player.cards))
	fmt.Println(player.getCard(0))
	fmt.Println(listToString(player.cards))

	deck := createDeck()
	deck.shuffle()
	deck.lift()
	fmt.Println(deck.peek(-1))
	fmt.Println(deck.peek(-1) == deck.take(-1))
}
