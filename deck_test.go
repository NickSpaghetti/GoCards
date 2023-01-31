package main

import (
	"os"
	"testing"
)

func TestNewDeck_HasFiftyTwoCards(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}
}

func TestNewDeck_FirstCard_IsAceOfSpades(t *testing.T) {
	d := newDeck()
	firstCard := d[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %v", firstCard)
	}
}

func TestNewDeck_LastCard_IsAceOfSpades(t *testing.T) {
	d := newDeck()
	lastCard := d[len(d)-1]
	if lastCard != "King of Clubs" {
		t.Errorf("Expected the last card to be King of Clubs, but got %v", lastCard)
	}
}

func cleanUpSaveTestFile() {
	testFileName := "_testingDeck.csv"
	_, err := os.OpenFile(testFileName, 0, 0666)
	if err == nil {
		os.Remove(testFileName)
	}
}

func TestSaveToDeck_NewDeck(t *testing.T) {
	cleanUpSaveTestFile()
	deck := newDeck()
	deck.saveToFile("_testingDeck")

	loadedDeck := loadDeckFromFile("_testingDeck.csv")
	if len(loadedDeck) != 52 {
		t.Errorf("Expted 52 cards in the deck, but got %d", len(loadedDeck))
	}
	cleanUpSaveTestFile()
}
