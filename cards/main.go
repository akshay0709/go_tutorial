package main

func main() {
	// var card string = "Ace of Spades"
	//card := newCard()
	// card = "Five of Diamonds" // Only use the ":" for init, reinit can only use "="
	//fmt.Println(card)

	// cards := newDeck()

	//cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
