//panic and recover

package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("something went wrong while run, message :", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	runApp(true)
}
