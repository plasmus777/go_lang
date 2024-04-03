package simple

import (
	"fmt"
	"strconv"
)

// Using the strconv package, it is possible to convert values from/to string values.
func Strconv() {
	// textBoolean := "treue"
	textBoolean := "true"
	boolean, error := strconv.ParseBool(textBoolean)
	// boolean, _ := strconv.ParseBool(textBoolean)
	fmt.Println("Boolean value: ", boolean)
	fmt.Println("Error: ", error)

	fmt.Println()

	textInteger := "-100"
	integer, error := strconv.ParseInt(textInteger, 10, 64)
	fmt.Println("Integer value: ", integer)
	fmt.Println("Error: ", error)

	fmt.Println()

	textFloat := "100.54"
	float, error := strconv.ParseFloat(textFloat, 64)
	fmt.Println("Float value: ", float)
	fmt.Println("Error: ", error)

}
