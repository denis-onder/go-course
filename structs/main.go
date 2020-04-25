package main

import "fmt"

type contactInfo struct {
	email       string
	phoneNumber int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	contact := contactInfo{email: "test@mail.com", phoneNumber: 15555355}

	p := person{
		firstName: "John",
		lastName:  "Doe",
		contact:   contact,
	}

	fmt.Printf("%+v\n", p)
}
