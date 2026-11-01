package main

import (
	"fmt";
	"learn-golang/fundamentals"
	"learn-golang/cards"
	"learn-golang/evenodd"
	"learn-golang/structs"
)

func main(){
	fmt.Println("Hello world!")
	fundamentals.Introduction()
	fundamentals.SwitchCase("linux")

	// Card Project
	// cards := []string{"Ace of Diamonds", cards.NewCard()} // declaring a slice
	cardDeck := cards.Deck{"Ace of Diamonds", cards.NewCard()}
	cardDeck = append(cardDeck, "Six of Spades") // appending to a slice
	// // does not modify original slice, creates a new slice and adds to it
	// for i, card := range cards { //index, current card := loop over slice
	// 	// we use := since we throw away the old index and element everytime we run through this loop
	// 	fmt.Println(i, card) // we need to use i, otherwise it shouldnt be declared
	// }
	cardDeck = cards.NewDeck()
	cardDeck.Print() // cardDeck is the receiver
	hand, remainingCards := cards.Deal(cardDeck, 5)
	hand.Print()
	remainingCards.Print()

	// cardDeck.SaveToFile("my_cards")
	// cardDeck2 := cards.NewDeckFromFile("my_cards")
	// cardDeck2.Print()

	// cardDeck.Shuffle()
	// cardDeck.Print()

	evenodd.IsEvenOrOdd()

	structs.CreatePerson()

}
