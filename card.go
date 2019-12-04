package main

import (
	"errors"
	"strings"
)

type card struct {
	// this will be passed around by reference
	// struct variables are pointers really:w

	suit string
	face string
}

func (c card) toString() string {
	return c.face + " of " + c.suit
}

func fromString(s string) (card, error) {
	data := strings.Split(s, " ")
	// the returned card is never nil, it may be just empty
	var c card
	if len(data) != 3 || data[1] != "of" {
		return c, errors.New("card string incorrect")
	}
	c.suit = data[2]
	c.face = data[0]
	return c, nil
}

// type X variables can be passed to functions as a receiver with *X receiver types, and it will be correct and &X will be passed.
