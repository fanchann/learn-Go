package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "ups!"
	}
}

func main() {
	var kosong interface{} = Ups(10)
	fmt.Println(kosong)
}
