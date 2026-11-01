package cards

import (
	"fmt"
	"strings"
	"os"
	"math/rand"
	"time"
)

type Deck []string

func NewDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // _ to ignore the index since we don't use it
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards

}

// d is a receiver and this is a "Receiver function"
// receiver is like this/self in Java/Python terms
func (d Deck) Print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func Deal(d Deck, handSize int) (Deck, Deck){ // deck divided into 2 slices
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

// byte slice is an array of ascii codes of the characters of a string
func (d Deck) SaveToFile(filename string) error{
	return os.WriteFile(filename, []byte(d.toString()), 0666) // anyone can read and write the file
}

func NewDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs),",")
	return Deck(s)
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // random int64 number with time
	r := rand.New(source) // type rand that takes in a seed source

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}