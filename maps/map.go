package maps

// Map needs to have same type on the keys side and same type on the values side
// It is a dictionary essentially

import (
	"fmt"
)

func CreateMap() {
	var colors map[string]string // map[]
	colors = make(map[string]string) // same as above

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors["white"] = "#ffffff" // we use ["value"] which needs to make sure its typed

	// delete(colors, "white")

	PrintMap(colors)
}

func PrintMap(c map[string]string) {
	for color, hex := range c{
		fmt.Println("Hex code for", color, "is", hex)
	}
}