package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected: 52, Size: %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected: Ace of Spades, Recieved: %s", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected: King of Clubs, Recieved: %s", d[0])
	}
}
func TestFileOperations(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	f := readFromFile("_deckTesting")
	if len(f) != 52 {
		t.Errorf("Expected: 52, Size: %d", len(f))
	}
	os.Remove("_deckTesting")
}
