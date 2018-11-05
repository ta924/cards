package main

func main()  {

	//var = create a new variable
	// card is the variable name
	//string would be the variable type
	// = "Ace of Spades" value assigned to the variable card
	//Create a new variable and assign it a default value
	//Type can be omitted and inferred by the runtime
	//var card string = "Ace of Spades"
	//var card = "Ace of Spades"
	//short hand approach of the same reference, := tells the compiler I want to create a type and have it inferred.
	//:= can only be used for variable creation
	//card := "Ace of Spades"

	//initialization via method
	//card := newCard()

	//creates a slice of strings
	//Go supports arrays and slices, arrays are fixed length and slices can grow and shrink on demand
	//cards := []string{newCard(), "Ace of Spades"}

	//Using custom type deck
	//cards := deck{newCard(), "Ace of Spades"}

	//append doesn't modify the original dataset, instead returns a new collection -- immutable
	//cards = append(cards, "Six of Spades")

	//cards := newDeck()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()

	//fmt.Println("----------------------------")

	//remainingCards.print()

	//Iterate over a closed set, range iterates over every record
	//i = index
	//card = current card in the range
	/*for i, card := range cards  {
		fmt.Println(i, card)
	}*/

	//cards.print()

	//Basic Go Types
	//bool, string, int, float64, etc.

	//print out the variable
	//fmt.Println(cards)


	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")

	//cards.print()


	cards := newDeck()

	cards.shuffle()

	cards.print()


}

//return type is specified after the method signature
/*func newCard() string {
	return "Five of Diamonds"
}*/
