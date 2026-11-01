package cards

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %s", d[0])
	}

	if d[len(d) - 1] != "Four of Diamonds" {
		t.Errorf("Expected last card to be Four of Diamonds, but got %s", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := NewDeck()
	deck.SaveToFile("_decktesting")

	loadedDeck := NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %d", len(loadedDeck))
	}
	
	os.Remove("_decktesting")
}