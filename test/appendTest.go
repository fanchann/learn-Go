package main

import (
	"fmt"
)

func main() {
	// names := [5]string{
	// 	"Fanny",
	// 	"Arif",
	// 	"Nasrudin",
	// 	"Farda",
	// 	"Ayu",
	// }
	// newName := names[2:4]
	// sliceOutput := append(newName, "kuntul")
	// fmt.Println(sliceOutput)
	// fmt.Println(names)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Farda"
	newSlice[1] = "Ayu"

	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{
		1, 2, 3, 4, 5,
	}
	iniSlice := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
