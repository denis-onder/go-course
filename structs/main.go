package main

import "fmt"

type contactInfo struct {
	email       string
	phoneNumber int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) greet() {
	fmt.Printf("Hello! My name is %s %s\n", p.firstName, p.lastName)
}

func main() {
	contact := contactInfo{email: "test@mail.com", phoneNumber: 15555355}

	p := person{
		firstName:   "John",
		lastName:    "Doe",
		contactInfo: contact,
	}

	fmt.Printf("%+v\n", p)

	p.updateName("Jim")
	p.greet()
}
