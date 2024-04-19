package exercises

import "fmt"

//slice = "2, 8, 3, 10, 5, 4, 7, 9, 1"
//Given the slice above, ranging from 1 to 10, create two variables, the first with the sum of all numbers ranging from 1 to 5 and the second with the sum of all numbers ranging from 6 to 10.
//Print the results.
func Ex2() {
	slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	var (
		sum1 int
		sum2 int
	)

	for i := 0; i < len(slice); i++ {
		if slice[i] <= 5 {
			sum1 += slice[i]
		} else {
			sum2 += slice[i]
		}
	}

	fmt.Println("Sum 1:", sum1)
	fmt.Println("Sum 2:", sum2)
}
