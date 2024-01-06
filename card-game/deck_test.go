package main

import "testing"

func TestCreateDeck(t *testing.T) {
	var deck = createDeck()

	var cardsCount = len(deck)

	if cardsCount != 13*4 {
		t.Errorf("Expected deck length of %v, but got %v", 13*4, cardsCount)
	}

	if deck[0] != "One of Spades" {
		t.Errorf("Expected first element to be %v, but got %v", "One of Spades", deck[0])
	}

	if deck[len(deck)-1] == "King of Diamond" {
		t.Errorf("Expected first element to be %v, but got %v", "King of Diamond", deck[len(deck)-1])
	}

}
