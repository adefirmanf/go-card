package main

import (
	"log"
	"math/rand"
)

// Create a deck
func createDeck() []string {
	log.Println("Creating a deck...")
	values := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
	}
	log.Println(values)
	return values
}

// Shuffling cards
func shuffleCard(deck []string) []string {
	log.Println("Shuffling cards...")

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

func main() {
	shuffleCard(createDeck())
}
