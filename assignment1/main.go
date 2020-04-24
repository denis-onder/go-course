package main

import "fmt"

func printResult(value int, result string) {
	fmt.Printf("%d is %s\n", value, result)
}

func main() {
	// Create a slice of ints
	var slice []int
	// Populate slice with values from 0 to 10
	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}
	// Loop over the slice, and check if the value is either even or odd
	for _, value := range slice {
		if value%2 == 0 {
			printResult(value, "even")
		} else {
			printResult(value, "odd")
		}
	}
}
