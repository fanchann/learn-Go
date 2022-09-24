package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
	IsMarried     bool
}

func main() {

	fanny := Customer{
		Name:      "Fanny",
		Address:   "Wonowoso",
		age:       19,
		IsMarried: false,
	}

	fmt.Println(fanny)
}
