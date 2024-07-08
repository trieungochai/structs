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
	contact   contactInfo
}

func main() {
	// declaring Structs
	haitn := person{
		firstName: "Trieu",
		lastName:  "Hai",
		contact: contactInfo{
			email:   "haitn.ute@gmail.com",
			zipCode: 933,
		},
	}

	fmt.Printf("%+v", haitn)
}
