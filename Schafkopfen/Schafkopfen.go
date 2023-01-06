package Schafkopfen

import (
	"CardGames/CardDeck"
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Schafkopfen struct {
	Turn    int
	Deck    CardDeck.Deck
	Players []*CardDeck.Player
}

func (schafkopfen *Schafkopfen) writeOutputAll(message string) {
	fmt.Println(message)
	for _, player := range schafkopfen.Players {
		player.Stdout <- message
	}
}

func CreateSchafkopfen(with_players bool) Schafkopfen {
	schafkopfen := Schafkopfen{0, CardDeck.CreateDeck(CardDeck.SIX), []*CardDeck.Player{}}
	if with_players {
		for n := 0; n < 4; n++ {
			player := (CardDeck.CreatePlayer(string(rune(n))))
			schafkopfen.Players = append(schafkopfen.Players, &player)
		}
	}
	return schafkopfen
}

type IntCard struct {
	n    int
	card CardDeck.Card
}

type Art int8

const (
	RAMSCH Art = iota
	SAUSPIEL
	GEIER
	WENZ
	SOLO
)

func (art Art) String() string {
	switch art {
	case GEIER:
		return "Geier"
	case WENZ:
		return "Wenz"
	case SOLO:
		return "Solo"
	case SAUSPIEL:
		return "Sauspiel"
	case RAMSCH:
		return "Ramsch"
	}
	return ""
}

type Spielweise struct {
	spieler []*CardDeck.Player
	farbe   CardDeck.Color
	art     Art
}

func findWinner(cards []CardDeck.Card, spielweise Spielweise) int {
	// Ober Check
	if spielweise.art != WENZ {
		for color := CardDeck.EICHEL; color >= CardDeck.SCHELLE; color-- {
			for n, card := range cards {
				if card.GetValue() == CardDeck.OBER && card.GetColor() == color {
					return n
				}
			}
		}
	}

	// Unter Check
	if spielweise.art != GEIER {
		for color := CardDeck.EICHEL; color >= CardDeck.SCHELLE; color-- {
			for n, card := range cards {
				if card.GetValue() == CardDeck.UNTER && card.GetColor() == color {
					return n
				}
			}
		}
	}

	var temp_card IntCard

	// Höhere Farbe Check
	if spielweise.art != GEIER && spielweise.art != WENZ {
		farbe := CardDeck.HERZ
		if spielweise.art == SOLO {
			farbe = spielweise.farbe
		}
		for n, card := range cards {
			if card.GetColor() == farbe && (card.GetValue() > temp_card.card.GetValue() || temp_card.card.GetValue() == 0) {
				temp_card = IntCard{n, card}
			}
		}
	}
	if temp_card.card.GetColor().String() != "" {
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

/*
Beispieleingaben:

	Geier
	Wenz
	Solo:Herz, Solo:Eichel
	Sauspiel:Herz, Sauspiel:Eichel
	Ramsch
*/
func requestSpielweise(player *CardDeck.Player) (Art, CardDeck.Color) {
	player.Stdout <- "Was würdest du spielen?"
	for {
		response := strings.Trim(<-player.Stdin, " \t\n\r")
		parts := strings.Split(response, ":")
		for spielart := RAMSCH; spielart <= SOLO; spielart++ {
			if strings.EqualFold(spielart.String(), parts[0]) {
				if spielart == SOLO || spielart == SAUSPIEL {
					for color := CardDeck.EICHEL; color >= CardDeck.SCHELLE; color-- {
						if strings.EqualFold(color.String(), parts[1]) {
							return spielart, color
						}
					}
				}
				return spielart, 0
			}
		}
	}
}

func (schafkopfen *Schafkopfen) RunRound() {
	schafkopfen.Deck.Shuffle()
	schafkopfen.Deck.Lift()
	schafkopfen.writeOutputAll("Die Runde hat begonnen")

	// Karten austeilen
	for i := 0; i < 2; i++ {
		for _, player := range schafkopfen.Players {
			for n := 0; n < 4; n++ {
				player.AddCard(schafkopfen.Deck.Take(0))
			}
		}
	}

	// Spielweise herausfinden
	var weise Spielweise = Spielweise{art: RAMSCH}
	for _, player := range schafkopfen.Players {
		art, farbe := requestSpielweise(player)
		if art > weise.art {
			weise.art = art
			weise.farbe = farbe
			weise.spieler = []*CardDeck.Player{player}
		} else if art == weise.art {
			if farbe > weise.farbe {
				weise.farbe = farbe
				weise.spieler = []*CardDeck.Player{player}
			}
		}
	}
	schafkopfen.writeOutputAll("Gespielt wird: " + weise.art.String())

	// Beim Sauspiel den zweiten Spieler zu den spielenden hinzufügen
	if weise.art == SAUSPIEL {
		var gesuchte_sau = CardDeck.CreateCard(weise.farbe, CardDeck.ASS)
		for _, player := range schafkopfen.Players {
			if slices.Contains(player.ListCards(), gesuchte_sau) {
				weise.spieler = append(weise.spieler, player)
			}
		}
	}

	// Spielrunden
	beginner := schafkopfen.Turn + 1

	for n := 0; n < 8; n++ {
		gelegte_karten := []CardDeck.Card{}

		// jeweils die gelegten Karten auswählen
		for i := 0; i < len(schafkopfen.Players); i++ {
			player := schafkopfen.Players[(schafkopfen.Turn+i+beginner)%len(schafkopfen.Players)]
			player.Stdout <- "Zu legende Karte:"
			for {
				number, err := strconv.Atoi(strings.Trim(<-player.Stdin, " \t\n\r"))
				if number >= len(player.ListCards()) {
					player.Stdout <- "Invalide Karte"
					continue
				}
				if err == nil {
					gelegte_karten = append(gelegte_karten, player.GetCard(int8(number)))
					schafkopfen.writeOutputAll(gelegte_karten[i].String())
					break
				}
			}
		}
		winner := findWinner(gelegte_karten, weise)
		gewinner := schafkopfen.Players[(beginner+winner)%len(schafkopfen.Players)]

		schafkopfen.writeOutputAll("Gewonnen hat: " + gelegte_karten[winner].String())

		for _, card := range gelegte_karten {
			gewinner.Punktzahl += int8(math.Max(float64(card.GetValue()), 0))
		}
		beginner = beginner + winner
	}

	// Punktestand ausgeben
	if weise.art == RAMSCH {
		schafkopfen.writeOutputAll(fmt.Sprintf("Punktzahlen: %d %d %d %d", schafkopfen.Players[0].Punktzahl, schafkopfen.Players[0].Punktzahl, schafkopfen.Players[0].Punktzahl, schafkopfen.Players[0].Punktzahl))
	} else if weise.art == SAUSPIEL {
		var pz int8 = 0
		for _, player := range weise.spieler {
			pz += player.Punktzahl
		}
		schafkopfen.writeOutputAll(fmt.Sprintf("Die Spieler konnten %d Punkte erreichen", pz))
	} else if weise.art == SOLO || weise.art == WENZ || weise.art == GEIER {
		schafkopfen.writeOutputAll(fmt.Sprintf("Der Spieler konnte %d Punkte erreichen", weise.spieler[0].Punktzahl))
	}

	schafkopfen.Turn += 1
	schafkopfen.Turn %= len(schafkopfen.Players)
}
