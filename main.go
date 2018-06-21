package main

func main() {
	// Creating a new deck
	// cards := newDeck()

	// Dealing the cards
	// as deal() returns two decks, we need to assign them both
	// hand will be assigned to the first deck thats returned, remainingCards will be assigned the second return value
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// Saving cards to file
	// cards.saveToFile("my_cards")

	// Loading cards from file
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// Shuffling a deck
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
