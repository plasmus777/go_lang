package fluxcontrol

import "fmt"

//There is no "while" function in go; it is possible to use "for" as a replacement.
//for [condition == true] {
func For() {

	text := "Testing the for loop..."

	for i := 0; i < len(text); i++ {
		// if text[i] == 'r' {
		// 	continue
		// }

		fmt.Println(i, " -> ", string(text[i]))

		// if text[i] == 'r' {
		// 	break
		// }
	}

	fmt.Println("")
	fmt.Println("-- Table --")
	for i := 1; i <= 10; i++ {
		fmt.Println("----------")
		fmt.Println("-----", i, "-----")
		for j := 1; j <= 10; j++ {
			fmt.Println(i, " * ", j, " = ", i*j)
		}
		fmt.Println("----------")
	}
}
