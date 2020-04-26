package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://duckduckgo.com/")
	if err != nil {
		fmt.Printf("An error has occured! %v\n", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	res.Body.Read(bs)
	defer res.Body.Close()
	fmt.Println(string(bs))
}
