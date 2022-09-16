package main

import "fmt"

func main() {
	type name string
	type married bool

	var nama name = "Fanny"
	var isMarried married = false

	fmt.Println(nama)
	fmt.Println(isMarried)
}
