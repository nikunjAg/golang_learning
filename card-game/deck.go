package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Created a new type deck which extends
// its functionalities from a string slice
type deck []string

// Added a receiver for deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println("Index: ", i, " Card: ", card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveDeckToFile(filename string) error {
	return writeToFile(filename, d.toString())
}

func (d deck) shuffle() {

	// Using standard math/rand library
	/* rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	}) */

	// Custom Logic
	// Generates the current unix timestamp in nanoseconds
	// Pass it as a seed to create a source for a new Random function
	source := rand.NewSource(time.Now().UnixNano())
	randomNumberGen := rand.New(source)

	for i := 0; i < len(d); i++ {
		var j = randomNumberGen.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
}

// Utility Methods
func createDeck() deck {
	var cardSuits = []string{"Spades", "Ace", "Hearts", "Diamond"}
	var cardValues = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	var cards = deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func getSlice(d deck, start int, end int) (deck, deck) {
	return d[start:end], append(d[:start], d[end:]...)
}

func readDeckFromFile(filename string) (deck, error) {

	var data, err = readFile(filename)
	return deck(strings.Split(data, ", ")), err
}
