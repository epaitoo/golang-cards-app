package main

import (
	"os"
	"io/ioutil"
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

// Save to File Function
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


// Read from a File
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	} 

	// convert byte to a string type
	s := strings.Split(string(bs), ",")
	
	// change and return the type of string to a deck typre
	return deck(s)

}