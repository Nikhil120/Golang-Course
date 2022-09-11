package main

func main() {

	cards := NewDeckFromFile("my_cards")
	cards.Print()

	cards.Shuffle()
	cards.Print()
}
