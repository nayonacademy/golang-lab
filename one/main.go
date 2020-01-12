package main

func main() {
	// cards := deck{"one", "two", "three"}
	// cards = append(cards, "four")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards := newDeck()
	hand, remainsCard := deal(cards, 2)
	hand.print()
	remainsCard.print()
	// cards.print()
	// cards.saveToFile("mycards")
	cards.shuffle()
	cards.print()
}
