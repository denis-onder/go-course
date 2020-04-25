package main

import "fmt"

func main() {
	// colors := make(map[string]string) => Creating maps with the make function
	colors := map[string]string{
		"red":   "FF0000",
		"green": "00FF00",
		"blue":  "0000FF",
	}
	colors["white"] = "FFFFFF" // Add pair to map
	delete(colors, "blue")     // Remove pair from map via key
	fmt.Println(colors)
	// Iterate over map
	for k, v := range colors {
		fmt.Printf("%s => %s\n", k, v)
	}
}
