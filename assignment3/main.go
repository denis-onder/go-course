package main

import (
	"fmt"
	"io"
	"os"
)

type customWriter struct{}

func (customWriter) Write(buf []byte) (int, error) {
	bw, err := fmt.Println(string(buf))
	if err != nil {
		return 0, err
	}
	return bw, nil
}

func main() {
	w := customWriter{}    // Custom writer
	filename := os.Args[1] // Get filename as arg from the terminal
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("An error has occured:", err)
		os.Exit(1)
	}
	io.Copy(w, file)
}
