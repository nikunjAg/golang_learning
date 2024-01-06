package main

import (
	"fmt"
	"os"
)

func getDeck() deck {
	var filename = "data.txt"
	var deck, err = readDeckFromFile(filename)

	if err != nil {
		fmt.Println("Unable to read deck from the file")
	} else if deck != nil {
		return deck
	}

	fmt.Println("Creating a new Deck")
	deck = createDeck()

	return deck
}

func main() {

	var deck = getDeck()
	deck.print()

	// Perform some operations...
	// ...
	deck.shuffle()
	var _, remDeck = getSlice(deck, 2, 5)

	var err = remDeck.saveDeckToFile("data.txt")
	if err != nil {
		fmt.Println("Unable to save the deck to the file")
		os.Exit(1)
	}

	fmt.Println("Successfully saved the deck")
}

func getCardName() string {
	var card string = "Ace Of Spades"
	return card
}
