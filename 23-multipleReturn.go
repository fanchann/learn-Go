package main

import "fmt"

func sayHello() (string, string) {
	return "Fanny", "Arif"
}

func main() {
	firstName, _ := sayHello()
	fmt.Println(firstName)
}
