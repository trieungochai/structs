package main

import "fmt"

// defining Structs
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// declaring Structs
	haitn := person{
		firstName: "Trieu",
		lastName:  "Hai",
		contactInfo: contactInfo{
			email:   "haitn.ute@gmail.com",
			zipCode: 933,
		},
	}

	haitn.updateFirstName("趙")
	haitn.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
