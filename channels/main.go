package main

import (
	"fmt"
	"net/http"
)

func getStatus(website string) {
	res, err := http.Get(website)
	if err != nil {
		fmt.Printf("%s => Not available.\n", website)
	}
	fmt.Printf("%s => %s\n", website, res.Status)
}

func main() {
	sites := []string{"https://google.com/", "https://duckduckgo.com/", "https://golang.org/"}
	for _, site := range sites {
		getStatus(site)
	}
}
