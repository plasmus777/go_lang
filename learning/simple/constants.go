package simple

import "fmt"

func Constants() {
	const x = 5
	const y = 7

	fmt.Println("X = ", x, ", Y = ", y)
	fmt.Println()
	fmt.Println("Since the values are assigned as constants, they cannot be altered, otherwise an error will be launched.")
}
