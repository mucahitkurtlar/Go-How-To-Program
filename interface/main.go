package main

import (
	"fmt"
)

type iName interface {
	getFullName() string
}

type animal struct {
	genus   string
	species string
}

func (a animal) getFullName() string {
	return a.genus + " " + a.species
}

type person struct {
	name     string
	lastname string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastname
}

func printFullName(in iName) {
	fmt.Println(in.getFullName())
}

func main() {
	myPerson := person{name: "Mucahit", lastname: "KURTLAR"}
	myAnimal := animal{genus: "Homo", species: "sapiens"}
	printFullName(myPerson)
	printFullName(myAnimal)
}
