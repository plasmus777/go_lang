package exercises

import "fmt"

//Create an integer array with size = 2. Store the total sum of the array's values in a variable, and then print it using the console.
func Ex1() {
	var array = [2]int{1, 5}

	var sum int

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	fmt.Println("Total sum:", sum)
}
