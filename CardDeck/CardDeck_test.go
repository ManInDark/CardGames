package CardDeck

import (
	"testing"
)

func TestColorString(t *testing.T) {
	if !(EICHEL.String() == "Eichel") {
		t.Fatal("String function of Color not working properly")
	}
	if !(HERZ.String() == "Herz") {
		t.Fatal("String function of Color not working properly")
	}
	if !(SCHELLE.String() == "Schelle") {
		t.Fatal("String function of Color not working properly")
	}
	if !(BLATT.String() == "Blatt") {
		t.Fatal("String function of Color not working properly")
	}
}

func TestValueString(t *testing.T) {
	if !(NINE.String() == "Neun") {
		t.Fatal("String function of Value not working properly")
	}
	if !(SIX.String() == "Sechs") {
		t.Fatal("String function of Value not working properly")
	}
	if !(UNTER.String() == "Unter") {
		t.Fatal("String function of Value not working properly")
	}
	if !(ASS.String() == "Ass") {
		t.Fatal("String function of Value not working properly")
	}
	// Assume that everything works then
}

func TestCardCreation(t *testing.T) {
	card := Card{EICHEL, SIX}
	if !(card.color == EICHEL && card.value == SIX && card == CreateCard(EICHEL, SIX)) {
		t.Fatal("Card creation not working")
	}
}

func TestGetColor(t *testing.T) {
	card := Card{EICHEL, TEN}
	if !(card.GetColor() == EICHEL) {
		t.Fatal("Card GetColor not working")
	}
}
func TestGetValue(t *testing.T) {
	card := Card{EICHEL, TEN}
	if !(card.GetValue() == TEN) {
		t.Fatal("Card GetValue not working")
	}
}

func TestCardString(t *testing.T) {
	if !(Card{SCHELLE, EIGHT}.String() == "[Schelle Acht]") {
		t.Fatal("Card String not working")
	}
}

func TestCardListString(t *testing.T) {
	if !(ListToString([]Card{{HERZ, SEVEN}, {EICHEL, TEN}}) == "[[Herz Sieben], [Eichel Zehn]]") {
		t.Fatal("Card ListToString not working")
	}
}

func TestPlayerCreation(t *testing.T) {
	player := CreatePlayer("somename")
	if !(player.name == "somename") {
		t.Fatal("Player CreatePlayer not working")
	}
}
func TestPlayerAddGetCard(t *testing.T) {
	player := CreatePlayer("somename")
	player.AddCard(Card{EICHEL, OBER})
	if !(player.GetCard(0) == Card{EICHEL, OBER}) {
		t.Fatal("Player AddGetCard not working")
	}
}

func TestPlayerString(t *testing.T) {
	player := CreatePlayer("somename")
	player.AddCard(Card{EICHEL, OBER})
	if !(player.String() == "somename: [[Eichel Ober]]") {
		t.Fatal("Player String not working")
	}
}

func TestPlayerListCards(t *testing.T) {
	player := CreatePlayer("somename")
	player.AddCard(Card{EICHEL, OBER})
	if !(len(player.ListCards()) == 1 && player.ListCards()[0] == Card{EICHEL, OBER}) {
		t.Fatal("Player ListCards not working")
	}
}

func TestDeckCreate(t *testing.T) {
	deck := CreateDeck(SIX)
	if !(len(deck.cards) == 32) {
		t.Fatal("Deck Length not working")
	}
}

func TestDeckLift(t *testing.T) {
	deck := CreateDeck(SIX, SEVEN, EIGHT, NINE, UNTER, OBER, KÖNIG)
	deck.Lift(2)
	compare_deck := []Card{{BLATT, TEN}, {EICHEL, TEN}, {SCHELLE, ASS}, {HERZ, ASS},
		{BLATT, ASS}, {EICHEL, ASS}, {SCHELLE, TEN}, {HERZ, TEN}}
	for n, card := range deck.cards {
		if card != compare_deck[n] {
			t.Fatal("Deck Lift not working")
		}
	}
}

func TestDeckPeek(t *testing.T) {
	deck := CreateDeck(SIX, SEVEN, EIGHT, NINE, TEN, UNTER, OBER, KÖNIG)
	if !(deck.Peek(-1) == Card{EICHEL, ASS}) {
		t.Fatal("Deck Peek not working")
	}
}

func TestDeckTake(t *testing.T) {
	deck := CreateDeck(SIX, SEVEN, EIGHT, NINE, TEN, UNTER, OBER, KÖNIG)
	if !(deck.Take(-1) == Card{EICHEL, ASS}) {
		t.Fatal("Deck Peek not working")
	}
}

func TestDeckString(t *testing.T) {
	deck := CreateDeck(SIX, SEVEN, EIGHT, NINE, TEN, UNTER, OBER, KÖNIG)
	if !(deck.String() == "[[Schelle Ass], [Herz Ass], [Blatt Ass], [Eichel Ass]]") {
		t.Fatal("Deck String not working")
	}
}
