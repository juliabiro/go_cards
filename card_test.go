package main

import "testing"

func TestFromString(t *testing.T) {
	c, err := fromString("alma of korte")
	if err != nil {
		t.Errorf("failed to read from string: %s", err)
	}
	if c.face != "alma" || c.suit != "korte" {
		t.Errorf("card incorrectly constructed from string 'alma of korte'; face is %s, suit is %s ", c.face, c.suit)
	}
}
func TestCard_ToString(t *testing.T) {
	c := card{suit: "Spades", face: "Two"}

	r := c.toString()
	if r != "Two of Spades" {
		t.Errorf("expected 'Two of Spades', but got %s", r)
	}
}
