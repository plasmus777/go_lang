package simple

import "fmt"

//& -> memory reference
//* -> pointer reference
func Pointers() {
	x := 5
	y := &x

	fmt.Println(&x, y)

	*y = 3

	fmt.Println(x, *y)

	showReference(x, *y)

	addThree(&x)

	fmt.Println(x, *y)
}

func showReference(x int, y int) {
	fmt.Println(&x, &y)
}

func addThree(x *int) {
	*x += 3
}
