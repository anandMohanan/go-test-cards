package main

func main() {
	cleanFile("cards")
	cards := newDeck()
	cards.shuffle()
	d, _ := deal(cards, 5)
	d.print()
	cards.saveToFile("cards")
	// hand, remainingCards := deal(cards, 6)
	// hand.print()
	// remainingCards.print()
	//fmt.Println(cards.toString())
	//cleanFile("cards")
	//cards.saveToFile("cards")

}
