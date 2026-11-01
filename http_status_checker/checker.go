package httpstatuschecker

import (
	"net/http"
	"fmt"
)

func PrintStatus() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	// serially
	for _, link := range links {
		go checkLink(link) // go keyword implements this in a separate go routine
	}

	// Go routine is the engine that executes code - one running program (process)

}

func checkLink(link string) {
	_, err := http.Get(link) // Blocking call - frozen on this line
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
