package simple

import "fmt"

func Operators(x int, y int) {
	sum := x + y
	sub := x - y
	mult := x * y
	div := float64(x) / float64(y)

	fmt.Println("Operation Values:")
	fmt.Println("X = ", x)
	fmt.Println("Y = ", y)
	fmt.Println("")
	fmt.Println("X + Y = ", sum)
	fmt.Println("X - Y = ", sub)
	fmt.Println("X * Y = ", mult)
	fmt.Println("X / Y = ", div)
}
