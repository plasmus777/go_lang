package simple

import (
	"fmt"
	"math"
)

// Go does not require the usage of "implements" or similar terms - the implementation is implicit.
type Geometry interface {
	Area() float64
}

func DisplayGeometryArea(g Geometry) {
	fmt.Println("Area:", g.Area())
}

type Rectangle struct {
	height, width float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func Interfaces() {
	r := Rectangle{
		height: 2,
		width:  4,
	}

	c := Circle{
		radius: 3,
	}

	DisplayGeometryArea(&r)
	DisplayGeometryArea(&c)

	fmt.Println()

	//The empty interface can have different types of value
	var slice []interface{}

	slice = append(slice, 10)
	slice = append(slice, 8.4)
	slice = append(slice, true)
	slice = append(slice, "test")

	fmt.Println(slice)

	for index, value := range slice {
		if val, ok := value.(string); ok {
			fmt.Println(val+" is a string; index =", index)
		} else {
			fmt.Println("index =", index)
		}
	}
}
