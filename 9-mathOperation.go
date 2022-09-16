package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := a + b
	fmt.Println(c)

	c += 10
	fmt.Println(c)

	c++
	fmt.Println(c)

	negative := -10
	positive := 10
	fmt.Println(negative)
	fmt.Println(positive)
}
