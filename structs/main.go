package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	// akshay := person{"Akshay", "Pawar"}

	// akshay := person{firstName: "Akshay", lastName: "Pawar"}
	// fmt.Println(akshay)

	// var akshay person

	// akshay.firstName = "Akshay"
	// akshay.lastName = "Pawar"

	// fmt.Printf("%+v", akshay)

	akshay := person{
		firstName: "Akshay",
		lastName:  "Pawar",
		contactInfo: contactInfo{
			email:   "akshay@gmail.com",
			zipCode: 94000,
		},
	}

	// akshayPointer := &akshay
	// akshayPointer.updateName("ak")
	// shortcut for pointers - directly reference the object and the reciver will know its a pointer
	akshay.updateName("ak")
	akshay.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
