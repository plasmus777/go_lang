package arrays

import "fmt"

//In practice, the code written in this file is related to slices - which differ from arrays both in performance/memory usage and in having a fixed size.
//An array can be defined as array := [size]int{elements}, where the size is already determined.
func Arrays() {
	//array := []int{0, 2, 4, 6, 8, 10}
	array := make([]int, 0)
	for i := 0; i < 6; i++ {
		array = append(array, i*2)
	}

	fmt.Println("Array:", array)
	fmt.Println("Array[3]:", array[3])
	fmt.Println("Array size:", len(array))

	//Use the memory address to work with pointers
	appendValue(&array, 13)

	fmt.Println()
	average(array)
	fmt.Println()

	subArrays_1()
	fmt.Println()

	subArrays_2()
	fmt.Println()
}

//It is important to use pointers when you want to alter variables that are external to the function
func appendValue(a *[]int, value int) {
	fmt.Println("\nAppending value \"", value, "\" to the array.")

	*a = append(*a, value)

	fmt.Println("Array:", *a)
	fmt.Println("Array size:", len(*a))
}

func average(a []int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		fmt.Print(" ", a[i])
		sum += a[i]
	}

	fmt.Println("\nAverage: ", float64(sum)/float64(len(a)))
}

func subArrays_1() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	var subArray = make([]int, 0)

	for i := 0; i < len(array); i++ {
		if array[i] < 5 {
			subArray = append(subArray, array[i])
		}
	}

	fmt.Println("Array where array[i] < 5:", subArray)
}

func subArrays_2() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	subArray1 := array[:3]
	fmt.Println("Array 1:", subArray1)

	subArray2 := array[4:]
	fmt.Println("Array 2:", subArray2)

	lastItem := array[len(array)-1:]
	fmt.Println("Array 3 - last item:", lastItem)
}
