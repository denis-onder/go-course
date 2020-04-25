package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "FF0000",
		"green": "00FF00",
		"blue":  "0000FF",
	}
	fmt.Println(colors)
	// Iterate over map
	for k, v := range colors {
		fmt.Printf("%s => %s\n", k, v)
	}
}
