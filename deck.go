package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create A New Deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// loop through both cardSuits & cardVlues (nested Loops)
	// "_" represents unused variables
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append "Ace of Spades" to cards
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// fucntion to print the x values of each card
// iterate over the cards slices using a for loop
func (d deck) showCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return two separate decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Helper Function to convert slices of Strings to Strings
func (d deck) toString() string {
	// use the Join func from stringS to join Slices with a comma(",")
	return strings.Join([]string(d), ",")

}
