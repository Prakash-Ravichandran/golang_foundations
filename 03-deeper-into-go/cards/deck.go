package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


func (nonShuffledCards deck) shuffle() deck {

	for index := range nonShuffledCards {
		randomCardIndex := rand.IntN(len(nonShuffledCards))
		nonShuffledCards[index], nonShuffledCards[randomCardIndex] = nonShuffledCards[randomCardIndex], nonShuffledCards[index]
	}

	return nonShuffledCards
}

func (s deck) toString() string {
	return strings.Join([]string(s), ",")
}

func stringToByte(s string) []byte {
	return []byte(s)
}

func (d deck) saveToFile(filename string) error {
	var data []byte = stringToByte(d.toString())
	return os.WriteFile(filename, data, 0666)
	// with type casting
	// return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func stringToStringArr(s string) []string {
	return strings.Split(s, ",")
}

func (d deck) readFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	s := stringToStringArr(string(data))
	return deck(s)
}
