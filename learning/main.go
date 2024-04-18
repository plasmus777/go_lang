package main

import (
	"fmt"

	"github.com/plasmus777/go_lang/learning/arrays"
	"github.com/plasmus777/go_lang/learning/exercises"
	"github.com/plasmus777/go_lang/learning/fluxcontrol"
	"github.com/plasmus777/go_lang/learning/simple"
)

func main() {

	exercises.Ex1()

	input := 1

	for input >= 1 {
		fmt.Println("---------------------------------")
		fmt.Println("Please select a function to execute:")
		fmt.Println("---------------------------------")
		fmt.Println("1 (default) - hello_world")
		fmt.Println("2 - variables")
		fmt.Println("3 - operators")
		fmt.Println("4 - strings")
		fmt.Println("5 - constants")
		fmt.Println("6 - variable convertion")
		fmt.Println("7 - strconv (string convertion)")

		fmt.Println("8 - if")
		fmt.Println("9 - for")

		fmt.Println("10 - arrays")
		fmt.Println("11 - maps")

		fmt.Println("")
		fmt.Println("-1 - exit")
		fmt.Println("---------------------------------")

		fmt.Scan(&input)

		switch input {
		case 1:
			simple.Hello_world()

		case 2:
			simple.Variables()

		case 3:
			simple.Operators(2, 3)

		case 4:
			simple.Strings()

		case 5:
			simple.Constants()

		case 6:
			simple.Convertion()

		case 7:
			simple.Strconv()

		case 8:
			fluxcontrol.If()

		case 9:
			fluxcontrol.For()

		case 10:
			arrays.Arrays()

		case 11:
			arrays.Maps()

		case -1:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid option.")
		}
		fmt.Println("---------------------------------")
		input = 1
	}
}
