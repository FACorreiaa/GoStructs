package main

import "fmt"

type contactInfo struct {
	city    string
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
	//dont
	//fernando := person{"Fernando", "Correia", 30, "Apulia"}

	//do
	var ps person

	ps.firstName = "Fernando"
	ps.lastName = "Correia"
	ps.age = 30
	ps.contactInfo.city = "Apulia"
	ps.contactInfo.email = "fernandocorreia316@gmail.com"
	ps.contactInfo.zipCode = 4740079

	//ps.updateName("Antonio")
	//still prints "Fernando"

	//updated code with pointers to update values

	//ps is a reference to the struct
	//with pPointer we acess the memory location
	// pPointer := &ps
	// pPointer.updateName("Antonio")
	// ps.print()

	ps.updateName("Antonio")
	ps.print()

	//values get updated with slices
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

//acess person value in memory

//person* is a type description
//pointerToPerson is a receiver to a type *person
//at this stage ps = pointerToPerson
//(*pointerToPerson) manipulate the value of the pointer we want to acess
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%v", p)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
