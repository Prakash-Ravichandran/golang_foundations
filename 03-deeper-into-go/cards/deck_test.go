package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards:= newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %s", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %s", cards[len(cards)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
   os.Remove("test_deck")

	cards := newDeck()

	cards.saveToFile("test_deck")


	loadedCards := cards.readFromFile("test_deck")

	if len(loadedCards) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedCards))
	}

	os.Remove("test_deck")
}