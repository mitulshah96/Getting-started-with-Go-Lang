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
}

func main() {
	mitul := person{
		firstName: "mitul",
		lastName:  "shah",
		contactInfo: contactInfo{
			email:   "mitulshah96@gmail.com",
			zipCode: 94000,
		},
	}

	// mitulPoniter := &mitul

	mitul.updateName("urvik")
	mitul.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
