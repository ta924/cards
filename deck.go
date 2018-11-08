package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"fmt"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {

	//Create an empty type of deck
	cards := deck{}

	//Create a slice for suites
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	//Create a slice of values
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//Usage of underscore tells Go we don't care about it or plan to use it
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//receiver function
//never use this or self, instead refer to value...in this example d
//Using a receiver to describe the input criteria of the method
//Any variable of type deck has access to the print method
//Convention is to use 1 letter arguments in receivers functions
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Returning multiple return values, both of type deck
func deal(d deck, handSize int) (deck, deck){

	//Slices are zero index based
	//d[indexstartinclueded:indexendexclusive]
	//leaving out the left value means start from the beginning
	//leaving out the right value means include everything to the end
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d),",")
}


func (d deck) saveToFile(filename string)  error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Write out the error
		fmt.Println("Error: ", err)
		//pass back a non-zero value to represent a non-successful return type
		os.Exit(1)
	}
	//cast byte slice to string
	//Use strings package to split the string on the comma delimeter and return a slice of strings
	s := strings.Split(string(bs), ",")

	return deck(s)

}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//swap wht is in d[i] with d[newPosition] and vice versa
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

