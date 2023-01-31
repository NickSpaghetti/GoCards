package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		nextPos := rand.Intn(len(d) - 1)
		d[i], d[nextPos] = d[nextPos], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	/*
		0 - place hodler
		6 - users can rw
		6 - groups can rw
		6 - other can rw
	*/
	fileNameWithExt := fmt.Sprintf("%v.csv", fileName)
	return os.WriteFile(fileNameWithExt, []byte(d.toString()), 0666)
}

func loadDeckFromFile(fileName string) deck {
	fileExt := filepath.Ext(fileName)
	if fileExt == "" {
		var stringBuilder strings.Builder
		stringBuilder.WriteString(fileName)
		stringBuilder.WriteString(".csv")
		fileName = stringBuilder.String()
	}

	byteSlice, error := os.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	csvCards := string(byteSlice)
	cards := strings.Split(csvCards, ",")
	return deck(cards)
}
