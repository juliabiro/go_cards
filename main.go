package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.shuffle()

	hand, remaining := deal(cards, 5)
	fmt.Println(remaining)

	hand.saveToFile("myfirstsave")

	newDeckFromFile("myfirstsave").print()
}
