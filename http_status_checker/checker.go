package httpstatuschecker

import (
	"fmt"
	"net/http"
	"time"
)

func PrintStatus() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	c := make(chan string) // make a channel

	// serially
	for _, link := range links {
		go checkLink(link, c) // go keyword implements this in a separate go routine
		// this is a child go routine
		// not the same as main go routine
	}

	// fmt.Println(<-c) // Waits for a message by main routine, blocking call again!
	// fmt.Println(<-c)
	// fmt.Println(<-c) // 3 for 3 links, 4 will hang the program

	// Better way from above:
	for range links { // i can also have i := 0; i < len(links), i++
		// fmt.Println(<-c)
	}

	// for { // because i want it to do infinite checks
	// 	go checkLink(<-c, c)
	// }

	// for l := range c {
	// 	go checkLink(l, c) // better than <-c as that would be ambiguous
	// }

	for l := range c {
		go func() { // a function literal - lambda function
			time.Sleep(5 * time.Second)
			checkLink(l, c) // before 1.22, you had to pass l as an argument since main and child go routine would point to same memory address
			// but after 1.22, they made sure that it is different memory address of the loop variable by default
		}()
	}

	// Go routine is the engine that executes code - one running program (process)
	// One CPU core -> Go Scheduler -> Go routine (multiple) -> OS thread (multiple)
	// Scheduler runs one routine until it finishes or makes a blocking call like an HTTP request
	// By default, Go uses 1 core
	// If we configure to multiple cores, the scheduler runs one thread on each logical core

	// Concurrency - Multiple threads executing code, if one thread breaks, another one is picked up and worked on
	// Parallelism - Multiple threads executed at same exact time with multiple CPU's

	// Main routine is not going to wait for child routines - hence, we have channels (orchestrator between main and child go routines)
	// Channel is typed (you need to define which type can be shared)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Blocking call - frozen on this line
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down I think"
		c <- link
		return
	}

	// Sending data with channels
	// channel <- 5 // Sends value 5 into channel
	// myNumber <- channel // Waits for value to be sent into channel, when we get one, assign to myNumber
	// fmt.Println( <- channel // Wait for value to be sent, log out immediately when we get one
	// c <- "Yes its up"
	c <- link
	fmt.Println(link, "is up!")
}
