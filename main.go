package main

import (
	"fmt"
	// "io"
	"learn-golang/cards"
	"learn-golang/evenodd"
	"learn-golang/fundamentals"
	httpstatuschecker "learn-golang/http_status_checker"
	"learn-golang/interfaces"
	"learn-golang/maps"
	"learn-golang/structs"
	// "os"
)

func main() {
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
	structs.PointerGotcha()

	maps.CreateMap()

	eb := interfaces.EnglishBot{}
	sb := interfaces.SpanishBot{}

	interfaces.PrintGreeting(eb)
	interfaces.PrintGreeting(sb)

	// interfaces.SendRequest()

	interfaces.Example()

	// fmt.Println(os.Args[1]) // gets the filename

	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// io.Copy(os.Stdout, f)

	httpstatuschecker.PrintStatus()
}
