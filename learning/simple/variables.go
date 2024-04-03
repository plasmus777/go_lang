package simple

import (
	"fmt"
	"math/rand"
)

func Variables() {

	var number int
	var decimal float64

	number = rand.Int()

	decimal = float64(number / 100000)

	text := "- This is a test to check variables using go."

	fmt.Println(number, decimal, text)
}
