package simple

import "fmt"

func Strings() {
	text1 := "I saw a blue bird"
	text2 := "on the tree today."

	concat := text1 + text2

	fmt.Println("Text 1: ", text1)
	fmt.Println("Text 2: ", text2)
	fmt.Println()
	fmt.Println("Concatenation: ", concat)
}
