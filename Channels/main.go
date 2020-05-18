package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// fmt.Println(links)

	c := make(chan string) //creation of channel of type string

	for _, link := range links {
		go checkLink(link, c)

	}
	// Receiving the data from child routine via channel

	for l := range c {
		// fmt.Println(<-c)
		//anonymous function or lamda
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		// go
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")

		c <- link //sending data to main routine via channel
		return
	}
	fmt.Println(link, " is up")

	c <- link //sending data to main routine via channel
}
