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

	alex.print()
	alex.updateName("Jim", "Carry")
	alex.print()
}

func (p person) print() {
	fmt.Printf("Name: %s %s\n", p.firstName, p.lastName)
	fmt.Printf("Email: %s\n", p.contact.email)
	fmt.Printf("Phone: %s\n", p.contact.phone)
}

func (pointerToPerson *person) updateName(firstName string, lastName string)  {
	(*pointerToPerson).firstName = firstName
	(*pointerToPerson).lastName = lastName
}