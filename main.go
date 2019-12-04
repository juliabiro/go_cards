package main

import (
//"fmt"
)

func main() {
	cards := NewDeck()
	cards.shuffle()
	// why does this actually work? because slices are basically pointers
	// so when a slice is passed by value, it is still modifiable
	// reference types: slice, map, channel, pointer, function

	hand, _ := Deal(cards, 5)

	//fmt.Println(remaining)

	hand.saveToFile("myfirstsave")

	NewDeckFromFile("myfirstsave").print()
}
