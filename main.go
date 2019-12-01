package main

import (
	"fmt"
)

func main() {
	cards := NewDeck()
	cards.Shuffle()

	hand, remaining := Deal(cards, 5)
	fmt.Println(remaining)

	hand.SaveToFile("myfirstsave")

	NewDeckFromFile("myfirstsave").print()
}
