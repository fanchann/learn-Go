package main

import (
	"fmt"
)

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fanny := Man{"Fanny"}
	fanny.Married()
	fmt.Println(fanny)
}
