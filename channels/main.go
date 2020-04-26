package main

import (
	"fmt"
	"net/http"
)

func getStatus(website string, c chan string) {
	res, err := http.Get(website)
	if err != nil {
		c <- website + " => Not available."
		return
	}
	c <- website + " => " + res.Status
}

func main() {
	c := make(chan string)
	sites := []string{"https://google.com/", "https://duckduckgo.com/", "https://golang.org/"}
	for _, site := range sites {
		go getStatus(site, c)
	}
	for i := 0; i < len(sites); i++ {
		fmt.Println(<-c)
	}
}
