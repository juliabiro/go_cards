package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new deck type
type deck []card //slice, not array

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardFaces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Prince", "Queen", "King"}

	for _, s := range cardSuits {
		for _, f := range cardFaces {
			cards = append(cards, card{suit: s, face: f}) // doesn't modify original slice, but returns a new one
		}
	}
	return cards
}

func (d deck) print() {
	// d is a receiver, similar to this or self
	// no classes to have functions; rather types get access to certain functions
	// the receiver is a reference, it is passed by value

	for _, card := range d { //index, value for iteration
		// under the hood it is reintializing the index and value variables, hence the :=
		fmt.Println(card.toString())
	}
}

func Deal(d deck, handSize int) (deck, deck) { //return 2 decks: we need the remaining one too
	// why is deck a parameter and not a reciever? because it is an input
	// it is copied
	// it doesn't change d
	//
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	var ret []string
	for _, card := range d {
		ret = append(ret, card.toString())
	}
	// using Join makes sure there are no trailing separators
	return strings.Join(ret, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
	// []byte("string") converts to byteslice
}

func NewDeckFromFile(filename string) deck {
	read, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(-1)
	}
	var d deck
	for _, s := range strings.Split(string(read), ",") {
		c, err := fromString(s)
		if err != nil {
			errors.New("incorrect card string: " + s)
			os.Exit(-1)
		}
		d = append(d, c)
	}

	return d
}

func (d deck) shuffle() {
	//true random
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	randomroof := len(d) - 1
	// Shuffle is not available before go 1.10
	for i, _ := range d {
		newPosition := r.Intn(randomroof)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
