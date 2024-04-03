package simple

import (
	"fmt"
	"math/rand"
	"reflect"
)

func Variables() {

	var number int
	var decimal float64

	number = rand.Int()

	decimal = float64(number / 100000)

	text := "This is a test to check variables using go."

	fmt.Println(reflect.TypeOf(number), " - ", number)
	fmt.Println(reflect.TypeOf(decimal), " - ", decimal)
	fmt.Println(reflect.TypeOf(text), " - ", text)
}
