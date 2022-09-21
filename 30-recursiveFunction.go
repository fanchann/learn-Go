package main

import "fmt"

func loopFactorial(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}

func recursiveFunction(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFunction(value-1)
	}
}

func main() {
	var number = 4
	recursive := recursiveFunction(number)
	fmt.Println(recursive)
	loop := loopFactorial(number)
	fmt.Println(loop)
}
