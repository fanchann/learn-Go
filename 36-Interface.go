package main

import "fmt"

type hashName interface {
	getName() string
}

func sayHi(hashName hashName) {
	fmt.Println("Hello", hashName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

func main() {
	person := Person{Name: "Fanny Arif"}
	sayHi(person)
}
