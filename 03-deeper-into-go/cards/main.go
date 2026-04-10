/**
Function declaration and return type.
**/

package main

import "fmt"

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

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
