package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("cards before shuffle")
	cards.shuffle()
	fmt.Println("cards after shuffle")
	cards.print()
}
