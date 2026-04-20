package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}


func main() {
    var alex person
    alex = person {
		firstName: "Alex",
		lastName: "Anderson",
		contact: contactInfo{
			email: "alex.anderson@example.com",
			phone: "123-456-7890",
		},
	}

	fmt.Printf("%v", alex)
}