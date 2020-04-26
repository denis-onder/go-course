package main

import (
	"fmt"
	"net/http"
)

func getStatus(website string, c chan string) {
	res, err := http.Get(website)
	if err != nil {
		fmt.Printf("%s => Down\n", website)
		c <- website
		return
	}
	fmt.Printf("%s => %s\n", website, res.Status)
	c <- website
}

func main() {
	c := make(chan string)
	sites := []string{"https://google.com/", "https://duckduckgo.com/", "https://golang.org/"}
	for _, site := range sites {
		go getStatus(site, c)
	}
	for {
		go getStatus(<-c, c)
	}
}
