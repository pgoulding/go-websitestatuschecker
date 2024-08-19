package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("------ Website Status Checker ------")

	// List of targets to fetch
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// make the channel
	c := make(chan string)

	//iterate through the links
	for _, link := range links {
		// staqrt a go routine with the link and the channel
		go checkLink(link, c)
	}

	// infinite loop for repeating the calls
	for l := range c {
		// function literal
		go func(link string) {
			// pause the execution for 5 seconds
			time.Sleep(5 * time.Second)
			// pass the return of the channel back to the function with another channel.
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// fetch the link
	_, err := http.Get(link)

	// check for error, or downed links
	if err != nil {
		fmt.Println(link, "might be down!")
		// pass the link back to the channel
		c <- link
		return
	}

	// Success message
	fmt.Println(link, "is up!")

	// pass the link back into the channel
	c <- link
}
