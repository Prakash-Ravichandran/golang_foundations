package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // add new card to the slice

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
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
