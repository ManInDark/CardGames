from CardDeck import *

kritische = [
    Card(Color.HERZ, Value.KÖNIG),
    Card(Color.SCHELLE, Value.SEVEN),
    Card(Color.EICHEL, Value.SEVEN)
]


class Watten:

    turn: int
    deck: Deck
    players: List[Player] = []
    player_count = 2

    def __init__(self, large_round=True):
        self.deck = Deck([Value.SIX])
        if large_round:
            self.player_count *= 2
        for n in range(self.player_count):
            self.players.append(Player(str(n)))
        self.turn = 0

    def writeOutputAll(self, message: str):
        for player in self.players:
            player.writeOutput(message)

    def startRound(self):
        self.deck.shuffle()

        # Abheben des Kartenstapels und Nehmen von Kritischen
        self.deck.lift()
        self.players[self.turn - 1].writeOutput(f"Abgehobene Karte: {str(self.deck.peek())}")
        if self.deck.peek() in kritische:
            self.players[self.turn - 1].writeOutput("Karte genommen")
            self.players[self.turn - 1].addCard(self.deck.take(-1))

        # restliche Karten austeilen
        for player in self.players:
            for _ in range(3 - len(player.listCards())):
                player.addCard(self.deck.take(0))
        for player in self.players:
            for _ in range(2):
                player.addCard(self.deck.take(0))
            player.writeOutput(f"Karten: {player.listCards()}")

        for player in self.players:
            print(len(player.listCards()))

        self.players[self.turn + 1].stdin.writeline("bube")  # entfernen
        self.players[self.turn].stdin.writeline("eichel")  # entfernen

        # Trumpf und Farbe ansagen
        self.players[self.turn + 1].writeOutput("Gewünschter Schlag ist: ")
        schlag = self.players[self.turn + 1].awaitInput(lambda inp: Value[inp.strip().upper()])
        print(schlag)
        self.players[self.turn].writeOutput("Gewünschte Farbe ist: ")
        farbe = self.players[self.turn].awaitInput(lambda inp: Color[inp.strip().upper()])
        print(farbe)

        # Spielrunden
        beginner = self.turn
        punkte = [0, 0]
        haube = Card(farbe, schlag)

        def punkteHinzufügen(gelegte_karten: List[Card], karte: Card):
            print(gelegte_karten, karte)
            punkte[(beginner + gelegte_karten.index(karte)) % self.player_count] += 1
        for n in range(5):
            self.writeOutputAll(f"Runde {n+1}")
            gelegte_karten: List[Card] = []

            # jeweils die gelegten karten auswählen
            for i in range(self.player_count):
                p = self.players[(beginner + i) % self.player_count]
                p.stdin.writeline("0")  # entfernen
                p.writeOutput(f"Zu legende Karte: (0-{4-n})")
                card = p.getCard(p.awaitInput(lambda inp: int(inp)))
                self.writeOutputAll(f"Karte wurde gelegt: {card}")
                gelegte_karten.append(card)

            # Gewinner der Runde feststellen
            for card in kritische:
                if card in gelegte_karten:
                    punkteHinzufügen(gelegte_karten, card)
                    gelegte_karten = []

            if haube in gelegte_karten:
                punkteHinzufügen(gelegte_karten, haube)
                continue

            for card in gelegte_karten:  # schlagtrümpfe
                if card.getValue() == schlag:
                    punkteHinzufügen(gelegte_karten, card)
                    gelegte_karten = []

            temp_card = None
            for card in gelegte_karten:
                if card.getColor() == farbe:
                    if temp_card is None or temp_card.getValue() < card.getValue():
                        temp_card = card
            if not temp_card is None:
                punkteHinzufügen(gelegte_karten, temp_card)
                continue

            for card in gelegte_karten:
                if temp_card is None or card.getColor() == temp_card.getColor() and card.getValue() > temp_card.getValue():
                    temp_card = card
            if not temp_card is None:
                punkteHinzufügen(gelegte_karten, temp_card)  # hier sollten alle vergeben sein
                continue
        print(punkte)
        self.turn += 1


watten = Watten(False)
watten.startRound()
