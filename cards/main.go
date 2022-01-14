package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()

	// remainingCards.print()

	// cards.saveToFile("test")

	// fmt.Println(newDeckFromFile("teste"))
}
