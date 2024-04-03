package main

import (
	"fmt"
	"learning/simple"
)

func main() {
	fmt.Println("Please select a file to execute:")

	fmt.Println("1 - hello_world")
	fmt.Println("2 - variables")

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		simple.Hello_world()

	case 2:
		simple.Variables()

	default:
		simple.Hello_world()
	}
}
