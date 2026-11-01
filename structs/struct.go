package structs

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo contactInfo
}

func CreatePerson() {
	// alex := person{"Alex", "Anderson"} // relies on order of fields, not so nice!
	// alex = person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var bob person // zero value is assigned: ""/0/false
	bob.firstName = "Bob"
	bob.lastName = "Thompson"
	bob = person{
		firstName: "Bob",
		lastName: "Thompson",
		contactInfo: contactInfo{
			email: "bob@gmail.com",
			zipCode: 94000,
		}, // every line must have a comma
	}

	bobPointer := &bob // give me memory address of the value this variable is pointing at
	bobPointer.updateName("bobby")
	
	bob = *&bob // note this
	bob.updateName("bobs") // this is just person but it can be used in *person or person later
	bob.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// * gives you the value that the address is pointing at
// *person is a type description - means we are working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// here, it is an operator that turns into an actual value that the pointer is referencing
	(*pointerToPerson).firstName = newFirstName // won't update
}

func PointerGotcha() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice) // updates it because its a slice - a reference type
	// other reference types: slices, maps, channels, pointers, functions
	// think about pointers when you have value types: int, float, string, bool, structs
}

func updateSlice(s []string) {
	s[0] = "Bye"
}