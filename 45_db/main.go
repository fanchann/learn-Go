package main

import (
	"fmt"
	"golang-dasar/45_db/db"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
