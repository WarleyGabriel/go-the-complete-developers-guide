package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {
	// me := person{firstName: "Warley", lastName: "Paixão"}
	// fmt.Println(me)

	// var warley person

	// warley.firstName = "Warley"
	// warley.age = 26
	// fmt.Println(warley)
	// fmt.Printf("%+v", warley)

	jim := person{firstName: "Jim", lastName: "Party", age: 33, contactInfo: contactInfo{email: "jim@mail.com", zipCode: 12345}}
	// jimPointer := &jim
	// jimPointer.updateFirstName("New")
	jim.updateFirstName("New")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
