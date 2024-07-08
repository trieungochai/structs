package main

import "fmt"

// defining Structs
type person struct {
	firstName string
	lastName  string
}

func main() {
	// declaring Structs
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
