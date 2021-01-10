package main

var deckSize int = 20

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
