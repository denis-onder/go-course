package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p := person{"John", "Doe"}
	fmt.Println("Hello, " + p.firstName + " " + p.lastName)
}
