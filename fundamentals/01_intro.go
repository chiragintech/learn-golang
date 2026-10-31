package fundamentals

import "fmt"

func Introduction() {
	fmt.Println("Hi! I am Chirag!")

	// Variables
	var x int
	var y = 20
	y = 30
	fmt.Println("Number is:", x, y) // space between x and y

	// Types - int, bool, string, float64, byte (0xFF - 255)
	z := 5 // type inference using "walrus" operator
	z = int(5) // type conversion
	fmt.Println("Number is:", z)

	// Single line comment
	/*
	Multi line comment
	*/

	// You cannot concatenate with int or float, it needs to be strings

	// Multiple variable declaration
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)

	// Constants (no :=)
	// Constants can be primitive types like strings, integers, booleans and floats. They cannot be more complex types like slices, maps and structs.
	const pi = 3.4
	// const fullName = firstName + " " + lastName // Possible since it can be compiled in compile time
	// const currentTime = time.Now() // Not possible!

	// fmt.Printf() - Prints a formatted string to standard out.
	// fmt.Sprintf() - Returns the formatted string
	// Formatting handles the "format verbs"
	// %v handles all values, %s for strings, %d for integers, %f/%.2f for floats, %t for booleans
	s := fmt.Sprintf("I am %v years old\n", 10)
	fmt.Print(s)

	const name = "Saul Goodman"
	const openRate = 30.54
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)
	fmt.Print(msg)

	height := 10
	if height > 6 { // No parenthesis in condition
    	fmt.Println("You are super tall!")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough!")
	}

	// An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if, else if, and else blocks. (Limits scope!)
	if length := 20; length < 10 {
    	fmt.Printf("Email must be at least 10 characters, is %d\n", length)
	}
}

func SwitchCase(os string) string {
	var creator string
	switch os {
	// break statement is not required at the end of a case to stop it from falling through to the next case. The break statement is implicit in Go.
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"
	
	// all three of these cases will set creator to "A Steve"
    case "macOS":
        fallthrough
    case "Mac OS X":
        fallthrough
    case "mac":
        creator = "A Steve"
	
	default:
		creator = "Unknown"
	}
	return creator
}