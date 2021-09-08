package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

type card struct {
	suit string
	value string
}

func (d deck) print() []string {
	var cards []string
	for _, card := range d {
		cards = append(cards, fmt.Sprintf("%s of %s", card.value, card.suit))
	}
	return cards
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, card{
				suit:  suit,
				value: value,
			})
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (c card) toString() string {
	return fmt.Sprintf("%s of %s", c.value, c.suit)
}

func (d deck) toString() string {
	return strings.Join(d.print(), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) (d deck) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newSlice := strings.Split(string(bs), ",")
	for _, newCard := range newSlice {
		d = append(d, card{
			suit: strings.Split(newCard, " of ")[1],
			value: strings.Split(newCard, " of ")[0],
		})
	}
	return d
}

func (d deck) shuffle() deck {
	// let's randomize the source of random number generation so that
	//  the deck is not shuffled the same way every time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}
