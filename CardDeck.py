from enum import Enum
from io import StringIO
from random import randint, shuffle
from typing import List


class StringBuffer(StringIO):

    def __init__(self):
        super().__init__()

    def write(self, __s: str) -> int:
        pos = self.tell()
        self.seek(len(self.getvalue()))
        super().write(__s)
        self.seek(pos)

    def writeline(self, line: str):
        self.write(f"{line}\n")


class Color(Enum):
    EICHEL = 0
    SCHELLE = 1
    HERZ = 2
    BLATT = 3

    def __gt__(self, other):
        if not isinstance(other, Color):
            raise ValueError
        return self.value > other.value

    def __lt__(self, other):
        if not isinstance(other, Color):
            raise ValueError
        return self.value < other.value

    def __eq__(self, other):
        if not isinstance(other, Color):
            raise ValueError
        return self.value == other.value

    def __ne__(self, other):
        if not isinstance(other, Color):
            raise ValueError
        return self.value != other.value


class Value(Enum):
    SIX = 6
    SEVEN = 7
    EIGHT = 8
    NINE = 9
    TEN = 10
    BUBE = 11
    DAME = 12
    KÖNIG = 13
    ASS = 14

    def __gt__(self, other):
        if not isinstance(other, Value):
            raise ValueError
        return self.value > other.value

    def __lt__(self, other):
        if not isinstance(other, Value):
            raise ValueError
        return self.value < other.value

    def __eq__(self, other):
        if not isinstance(other, Value):
            raise ValueError
        return self.value == other.value

    def __ne__(self, other):
        if not isinstance(other, Value):
            raise ValueError
        return self.value != other.value


class Card:
    color: Color
    value: Value

    def __init__(self, color: Color, value: Value) -> None:
        self.color = color
        self.value = value

    def __str__(self) -> str:
        return self.__repr__()

    def __repr__(self) -> str:
        return f"[{self.color}, {self.value}]"

    def getColor(self):
        return self.color

    def getValue(self):
        return self.value

    def __eq__(self, other):
        if not isinstance(other, Card):
            raise ValueError
        return self.value == other.value and self.color == other.color


class Player:

    cards: List[Card]
    name: str
    stdout: StringBuffer
    stdin: StringBuffer

    def __init__(self, name: str):
        self.name = name
        self.stdout = StringBuffer()
        self.stdin = StringBuffer()
        self.cards = []

    def addCard(self, card: Card):
        self.cards.append(card)

    def getCard(self, index: int) -> Card:
        return self.cards.pop(index)

    def listCards(self) -> List[Card]:
        return self.cards

    def writeOutput(self, message: str):
        """sends the player a message"""
        self.stdout.writeline(message)

    def readInput(self, length: int = -1) -> str:
        """reads wether there has been a textinput"""
        return self.stdin.readline(length)

    def awaitInput(self, validation) -> str:
        """Reads the user input until validation returns something"""
        while True:
            try:
                text = self.readInput()
                if len(text) > 0:
                    return validation(text)
            except:
                self.writeOutput("Invalid Input")
                continue


class Deck:

    cards: List[Card] = []

    def __init__(self, ignored_values: List[Value]):
        for color in Color:
            for value in Value:
                if value not in ignored_values:
                    self.cards.append(Card(color, value))

    def shuffle(self):
        shuffle(self.cards)

    def lift(self):
        """= Abheben """
        index = randint(0, len(self.cards))
        self.cards = self.cards[index:len(self.cards)] + self.cards[0:index]

    def peek(self, index: int = -1) -> Card:
        return self.cards[index]

    def take(self, index: int = 0) -> Card:
        return self.cards.pop(index)

    def __str__(self) -> str:
        return self.__repr__()

    def __repr__(self) -> str:
        return f"[{', '.join(str(card) for card in self.cards)}]"


if __name__ == "__main__":
    d = Deck([Value.SIX])
    d.shuffle()
    d.lift()
    print(d)
    print(d.peek())
