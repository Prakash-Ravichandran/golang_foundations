package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://golang.org",
		"http://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
	   go checkLink(<-c, c)
		// here go allows us to pass <- c as link because it could able to refer it as a string.
	}
}

func checkLink(link string, c chan string) {
		_, err := http.Get(link)
	 if err != nil {
		fmt.Println(link, "might be down");
		c <- link
		return
	 }
	 fmt.Println(link, "is up");
	 c <- link
}