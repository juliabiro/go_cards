package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d) != 52 {
		t.Errorf("expected deck length 52, but got %d", len(d))
	}

	if d[0].toString() != "Ace of Spades" {
		t.Errorf("expected new decks first card should be Ace of Spades, but got %s", d[0])
	}
	if d[len(d)-1].toString() != "King of Clubs" {
		t.Errorf("expected new decks last card should be King of Clubs, but got %s", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := NewDeck()
	hand, _ := Deal(d, 10)
	if len(hand) != 10 {
		t.Errorf("expected hand length 10, but got %d", len(hand))
	}
}
func TestNewDeckFromFile(t *testing.T) {
	d := NewDeckFromFile("test_hand.data")
	if d[0].toString() != "Two of Diamonds" {
		t.Errorf("expected test decks first card should be Two of Diamonds, but got %s", d[0])
	}
}
func TestDeck_Print(t *testing.T) {
}
func TestDeck_ToString(t *testing.T) {
}
func TestDeck_SaveToFile(t *testing.T) {
	testfilename := "_decktesting"

	os.Remove(testfilename)
	d := NewDeck()
	d.saveToFile(testfilename)

	if _, err := os.Stat(testfilename); os.IsNotExist(err) {
		t.Errorf("deck file didnt get generated")
	}

	os.Remove(testfilename)
}
func TestDeck_Shuffle(t *testing.T) {
	d := NewDeck()

	first := d[0]
	d.shuffle()
	if d[0] == first {
		t.Errorf("the shuffling didn't work, as the first card is uncahnged, %s", d[0])
	}
}
