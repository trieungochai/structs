package main

import "fmt"

// defining Structs
type person struct {
	firstName string
	lastName  string
}

func main() {
	// declaring Structs
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)
}
