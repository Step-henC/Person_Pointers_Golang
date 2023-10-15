package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	nickNames []string
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Johns",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
		nickNames: []string{"Buddy", "Peanut", "Big-Head"},
	}

	//	jimPointer := &jim // give me memory address // points at memory address of struct

	//jimPointer.updateFirstName("Jimmy")

	//do not need the above code, Golang will handler the pointer for you if you make a receiver of
	// type pointer to struct
	jim.updateFirstName("jimmy") // still works

	updateNickNames(jim.nickNames)
	jim.print()

}

//the star on a type is not an operator
//mainly says this operation can only be done on a pointer to a type 'person'
//this is not a person, but a reciever for a pointer TO a person (a diff type)
func (pointerToPerson *person) updateFirstName(newName string) {
	//p.firstName = newName
	// need pointers for memory address to update name otherwise...
	//only updating a 'copy' of that variable

	(*pointerToPerson).firstName = newName // Do not give me the memory address
	//Give me the VALUE this memory address is pointing at
}

func (p person) print() {
	fmt.Printf("%+v", p) // prints all the fields of the person
}

func updateNickNames(n []string) {

	//still works without pointers
	//Golang makes copy of slice
	//but also has a pointer to the head
	// passed by reference, not value
	//along with other non-primitive types like functions, maps, channels etc.
	n[0] = "Chicken"
}

//turn address to value with *address
//turn value to address with &value
