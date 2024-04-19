package simple

import (
	"fmt"
	"strconv"
)

// It is possible to have as many init() functions as you wish.
// func init() {
// 	printMessage("The init function will run before anything else, including the main() function, as soon as the package is imported.")
// }

func Functions() {
	printMessage("This is a test message.")

	printMessage(strconv.Itoa(sumTwoNumbers(3, 7)))

	sum, sub, mult, div := operations(3, 2)
	printMessage(strconv.Itoa(sum) + "; " + strconv.Itoa(sub) + "; " + strconv.Itoa(mult) + "; " + strconv.FormatFloat(div, 'f', -1, 64))
}

func printMessage(message string) {
	fmt.Println("**********************")
	fmt.Println(message)
	fmt.Println("**********************")
}

// returns one value
func sumTwoNumbers(num1 int, num2 int) int {
	return (num1 + num2)
}

// returns multiple values
func operations(num1 int, num2 int) (sum int, sub int, mult int, div float64) {
	sum = (num1 + num2)
	sub = (num1 - num2)
	mult = (num1 * num2)
	div = (float64(num1) / float64(num2))

	return
}
