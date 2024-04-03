package simple

import (
	"fmt"
	"reflect"
)

func Convertion() {
	var x int8 = 2
	var y int = 3

	var x2 int = int(x)

	fmt.Println("X = ", x, "(", reflect.TypeOf(x), "), Y = ", y, "(", reflect.TypeOf(y), ")")
	fmt.Println()
	fmt.Println("X converted to int -> ", "X = ", x2, "(", reflect.TypeOf(x2), ")")
}
