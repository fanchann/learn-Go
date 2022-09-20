package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, ttl := range number {
		total += ttl
	}
	return total
}

func main() {
	total := sumAll(10, 10)
	fmt.Println(total)

	//jika ingin menggunakan type data slice
	slice := []int{10, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
