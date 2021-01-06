package main

// import (
// 	"fmt"
// )



func main() {

	// Declare a Slice of Cards
	// cards := newDeck()

	// Assign the values from the return deal func to 'hand' & 'remainingCards'
	// Both variables are of type dexk
	// hand, remainingCards := deal(cards, 5)

	// hand.showCards()
	// remainingCards.showCards()

	// fmt.Println(cards.toString()) 

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_ards")

	cards.showCards()

}
