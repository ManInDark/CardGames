package CardDeck

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
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
	return "" // unreacheable
}

type Card struct {
	color Color
	value Value
}

func (card Card) GetColor() Color {
	return card.color
}

func (card Card) GetValue() Value {
	return card.value
}

func (card Card) String() string {
	return "[" + card.color.String() + " " + card.value.String() + "]"
}

func ListToString(cards []Card) string {
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
	Stdout chan string
	Stdin  chan string
}

func CreatePlayer(name string) Player {
	return Player{[]Card{}, name, make(chan string), make(chan string)}
}

func (player Player) String() string {
	return player.name + ": " + ListToString(player.cards)
}

func (player *Player) AddCard(card Card) {
	player.cards = append(player.cards, card)
}

func (player *Player) GetCard(index int8) Card {
	card := player.cards[index]
	shifted := player.cards[index+1:]
	player.cards = append(player.cards[:index], shifted...)
	return card
}

func (player Player) ListCards() []Card {
	return player.cards
}

type Deck struct {
	cards []Card
}

func CreateDeck(ignored_values ...Value) Deck {
	deck := Deck{[]Card{}}
	for value := SIX; value <= ASS; value++ {
		if !(slices.Contains(ignored_values, value)) {
			for color := EICHEL; color <= BLATT; color++ {
				deck.cards = append(deck.cards, Card{color, value})
			}
		}
	}
	return deck
}

func (deck Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })
}

func (deck *Deck) Lift(given_index ...int) {
	index := rand.Intn(len(deck.cards) - 1)
	if len(given_index) > 0 {
		index = given_index[0]
	}
	lifted := deck.cards[:index]
	deck.cards = append(deck.cards[index:len(deck.cards)], lifted...)
}

func (deck Deck) Peek(index int) Card {
	if index == -1 {
		index = len(deck.cards) - 1
	}
	return deck.cards[index]
}

func (deck *Deck) Take(index int) Card {
	if index == -1 {
		index = len(deck.cards) - 1
	}
	card := deck.cards[index]
	shifted := deck.cards[index+1:]
	deck.cards = append(deck.cards[:index], shifted...)
	return card
}

func (deck Deck) String() string {
	return ListToString(deck.cards)
}
