package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Printf("Current deck:\n %s\n", cards.toString())
	shuffled := cards.shuffle()
	fmt.Printf("Shuffled deck:\n %s\n", shuffled.toString())
}
