package main

import "fmt"

type checkUsername func(name string) bool

func CheckUser(name string, check checkUsername) {
	if check(name) {
		fmt.Println(name + " was registered please any")
	} else {
		fmt.Println(name + " welcome")
	}
}

func main() {
	checkIsAvailable := func(name string) bool {
		return name == "admin" || name == "fanny"
	}
	CheckUser("admin", checkIsAvailable)
	CheckUser("fanny", checkIsAvailable)

}
