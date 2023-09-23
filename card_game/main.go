package main

// we cannot assign a value to a variable outside the func, but we can decalre it
func main() {
	//var card string = "Ace of Spades"
	//:= is only used when assigning new variable
	cards := newDeckFromFile("my_cards")
	cards.print()

	//cards.print()
}

// files in same package can freely call functions defined in other files
