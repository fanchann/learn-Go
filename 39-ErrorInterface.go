package main

import (
	"errors"
	"fmt"
)

func Pembagi(number int, bagi int) (int, error) {
	if bagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return number / bagi, nil
	}
}

func main() {
	result, err := Pembagi(1000, 10)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", err)
	}
}
