package evenodd

import "fmt"

func IsEvenOrOdd() {
	list := make([]int, 11) // create [0,0,0...] with 11 elements
	for i := range list {
		list[i] = i
	}
	for _, value := range list {
		if value % 2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}