package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Error("Expected 52 cards but received", len(d))
	}

	c := d[0]
	if c.value != "Ace" && c.suit != "Hearts" {
		t.Error("Expected Ace of Hearts as first card but received", c.value, c.suit)
	}

	c2 := d[len(d) -1]
	if c2.value != "King" && c2.suit != "Clubs" {
		t.Error("Expected King of Clubs as last card but received", c2.value, c2.suit)
	}
}

func TestLoadFromDeckFile(t *testing.T) {
	testDeckFilename := "_testdeck.txt"
	err := os.Remove(testDeckFilename)
	if err != nil {
		t.Log("Problem while cleaning up test deck file", err)
	}

	d := newDeck()
	saveErr := d.saveToFile(testDeckFilename)
	if saveErr != nil {
		t.Error("Failed to save deck to file", saveErr)
	}

	newDeck := newDeckFromFile(testDeckFilename)
	if len(newDeck) != 52 {
		t.Errorf("Expected full deck from file but received %d cards", len(newDeck))
	}

	if d[0] != newDeck[0] {
		t.Errorf("Expected top card to be the same between d and newDeck but saw %s and %s", d[0], newDeck[0])
	}
}
