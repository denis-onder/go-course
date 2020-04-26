package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://duckduckgo.com/")
	if err != nil {
		fmt.Printf("An error has occured! %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	// buf := make([]byte, 32*1024)
	// res.Body.Read(buf)
	// fmt.Println(string(buf))
	io.Copy(os.Stdout, res.Body)
}
