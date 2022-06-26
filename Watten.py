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

    def startRound(self):
        self.deck.shuffle()

        # Abheben des Kartenstapels und Nehmen von Kritischen
        self.deck.lift()
        self.players[self.turn - 1].writeOutput(f"Abgehobene Karte: {str(self.deck.peek())}")
        if self.deck.peek() in kritische:
            self.players[self.turn - 1].writeOutput("Karte genommen")
            self.deck.take(-1)

        # restliche Karten austeilen
        for player in self.players:
            for _ in range(3 - len(player.listCards())):
                player.addCard(self.deck.take(0))
        for player in self.players:
            for _ in range(2):
                player.addCard(self.deck.take(0))
            player.writeOutput(f"Karten: {player.listCards()}")

        self.players[self.turn + 1].stdin.writeline("bube")
        self.players[self.turn].stdin.writeline("eichel")

        # Trumpf und Farbe ansagen
        self.players[self.turn + 1].writeOutput("Gewünschter Schlag ist: ")
        schlag = self.players[self.turn + 1].awaitInput(lambda inp: Value[inp.strip().upper()])
        print(schlag)
        self.players[self.turn].writeOutput("Gewünschte Farbe ist: ")
        farbe = self.players[self.turn].awaitInput(lambda inp: Color[inp.strip().upper()])
        print(farbe)


watten = Watten(False)
watten.startRound()
