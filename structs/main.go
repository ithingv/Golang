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

	person1 := person{
		firstName: "kim",
		lastName:  "hong",
		age:       32,
		contactInfo: contactInfo{
			email:   "ithing43@gmail.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v\n", person1)

	// zero value
	var person2 person
	fmt.Println(person2)
	fmt.Printf("%+v\n", person2)

	// not working
	// newName := "lee"
	// person1.updateName(&newName)
	leePointer := &person1
	leePointer.updateName("lee")
	person1.print()

	// shortcut
	lee.updateName("Lee")
	lee.print()
}

// not working
// func (p person) updateName(newFirstName *string) {
// 	p.firstName = *newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

/*
	pointer rule
	Turn address into value with *address
	TUrn value into address with &value
*/

func (p person) print() {
	fmt.Printf("%+v", p)
}
