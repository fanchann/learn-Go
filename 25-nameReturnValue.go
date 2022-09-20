package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Fanny Arif"
	lastName = "Nasrudin"
	return
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
