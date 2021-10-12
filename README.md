# go-playing-cards
Just messing around with a deck of cards

## Build
```bash
$ go build main.go deck.go
```

## Test
```bash
$ go test -v
```

## Use cases
* Create a new deck of cards with `newDeck()`: `cards := newDeck()`
* Create a new deck of cards from a CSV file
  * Save deck to file: `cards.saveToFile(nameOfFile)`
  * Load deck from file: `cards := newDeckFromFile(nameOfFile)`
* Shuffle cards: `shuffledCards := cards.shuffle()`
