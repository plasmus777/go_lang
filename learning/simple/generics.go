package simple

import "fmt"

func Generics() {
	slice1 := []int{5, 3, 7, 1}
	fmt.Println(slice1)
	fmt.Println(*reverse(slice1))

	fmt.Println()

	slice2 := []string{"b", "g", "a", "m"}
	fmt.Println(slice2)
	fmt.Println(*reverse(slice2))

}

//Creating custom constraints can make your code easier to read/understand.
type customConstraint interface {
	int | string
}

//Generics are useful when you want to reuse code dealing with different argument/return types.
//T is a generic type (in this case, either an int or a string).
func reverse[T customConstraint](slice []T) *[]T {
	reverseSlice := make([]T, len(slice))

	j := 0
	for i := len(slice) - 1; i >= 0; i-- {
		reverseSlice[j] = slice[i]
		j++
	}

	return &reverseSlice
}
