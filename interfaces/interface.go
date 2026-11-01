package interfaces

import(
	"fmt"
)

// Concrete types - struct, int, map, string
// Interface type - interface (bot)
// Interfaces are not generic types, so does not have them
// Interfaces are implicit, go takes care of the matching
// Interfaces are a contract to help us manage types, if the type's implementation is broke, it wont help
// Writing interfaces are tough
type Bot interface { // check all types and if you have a getgreeting that returns a string, you are also of type "bot"
	GetGreeting() string // method set
	// you can put input and return arguments
}

type EnglishBot struct {}
type SpanishBot struct {}

// we don't need to have eb as we are not using it
func (EnglishBot) GetGreeting() string {
	// custom logic
	return "Hi there"
}

func (SpanishBot) GetGreeting() string {
	return "Hola"
}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}

// func PrintGreeting(sb SpanishBot) {
// 	fmt.Println(sb.GetGreeting())
// }
