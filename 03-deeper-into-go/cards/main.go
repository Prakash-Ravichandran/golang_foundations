package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	// cards.print()

	// hands, remainingCards := deal(cards, 5)

	// hands.print()
	// remainingCards.print()

	cardsString := cards.toString()
	fmt.Println(cardsString)
	cards.saveToFile("my_cards")
}

/**
Function declaration and return type.
**/

//
// func main() {
// 	card := newCard()
// 	fmt.Println(card)
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }

// func main() {
// 	printState()
// }
