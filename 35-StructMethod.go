package main

import "fmt"

type Member struct {
	Name, Address string
	Age           int
}

func (member Member) sayHi(name string) {
	fmt.Println("Hello ", name, " My name is", member.Name)
}

func (greet Member) sayHu() {
	fmt.Println("Huuuu", greet.Name)
}

func main() {
	fanny := Member{
		Name:    "Fanny",
		Address: "Demak",
		Age:     19,
	}

	fanny.sayHi("eko")
	fanny.sayHu()
}
