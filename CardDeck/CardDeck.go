package CardDeck

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
)

type Color int8
type Value int8

const (
	EICHEL  Color = 4
	BLATT   Color = 3
	HERZ    Color = 2
	SCHELLE Color = 1
	SIX     Value = -4
	SEVEN   Value = -3
	EIGHT   Value = -2
	NINE    Value = -1
	UNTER   Value = 2
	OBER    Value = 3
	KÖNIG   Value = 4
	TEN     Value = 10
	ASS     Value = 11
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
	case UNTER:
		return "Unter"
	case OBER:
		return "Ober"
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

func CreateCard(color Color, value Value) Card {
	return Card{color, value}
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
	cards     []Card
	name      string
	Stdout    chan string
	Stdin     chan string
	Punktzahl int8
}

func CreatePlayer(name string) Player {
	return Player{[]Card{}, name, make(chan string), make(chan string), 0}
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
	accepteable_values := []Value{SIX, SEVEN, EIGHT, NINE, UNTER, OBER, KÖNIG, TEN, ASS}
	for value := SIX; value <= ASS; value++ {
		if !(slices.Contains(ignored_values, value)) && slices.Contains(accepteable_values, value) {
			for color := SCHELLE; color <= EICHEL; color++ {
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
	println(len(deck.cards) - 1)
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
