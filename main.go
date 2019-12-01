package main

import (
//"fmt"
)

func main() {
	cards := NewDeck()
	cards.shuffle()

	hand, _ := Deal(cards, 5)
	//fmt.Println(remaining)

	hand.saveToFile("myfirstsave")

	NewDeckFromFile("myfirstsave").print()
}
