package main

import (
	"os"
	"testing"
)

//All tests are called by the runner with *testing.T
//By convention all tests should start with a capital letter
func TestNewDeck(t *testing.T)  {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T)  {

	filename := "_decktesting"

	cleanUpFiles(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards from loaded deck, but got %v", len(loadedDeck))
	}

	cleanUpFiles(filename)
}

func cleanUpFiles(filename string)  {
	os.Remove(filename)
}
