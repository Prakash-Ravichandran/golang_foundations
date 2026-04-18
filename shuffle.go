// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	cards := deck()
	fmt.Println("unshuffled cards", cards)

	shuffledCards := shuffle(cards)
	fmt.Println("shuffled cards", shuffledCards)
}

func deck() []string {
	cards := []string{"Ace of spades", "Two of kings", "Three of Diamonds", "Four of Hearts"}
	return cards
}

func shuffle(deck []string) []string {

	nonShuffledCards := deck
	length := len(nonShuffledCards)
	min, max := 0, length-1

	for i, currentCard := range nonShuffledCards {
		randomCardIndex := rand.IntN(max - min + i)

		firstCard := currentCard
		randomCard := nonShuffledCards[randomCardIndex]
		nonShuffledCards[randomCardIndex] = firstCard
		nonShuffledCards[i] = randomCard
	}

	return nonShuffledCards
}
