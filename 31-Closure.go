package main

import "fmt"

func main() {
	counter := 0

	closure := func() {
		fmt.Println("counter++")
		counter++
	}

	closure()
	closure()
	fmt.Println(counter)
}
