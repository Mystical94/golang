package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type employee struct {
	person
	empId int
}

func singleStruct() {
	x := person{age: 24, firstName: "Shubham", lastName: "Gupta"}
	fmt.Println(x)
	fmt.Println(x.firstName)
}

func (e employee) inheritedStruct() {
	fmt.Println(e, " "+"I am a free person")
}

func main() {
	singleStruct()
	e1 := employee{person: person{"Shubham", "Gupta", 24}, empId: 3}
	e1.inheritedStruct()
}
