package simple

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Plate struct {
	letters string
	numbers string
}

type Car struct {
	name  string
	color string
	plate Plate
}

func Structs() {
	fmt.Println("Creating car...")

	car := Car{
		name:  "Honda Civic Type R",
		color: "red",
		plate: Plate{
			letters: "ABC",
			numbers: "1234",
		},
	}

	fmt.Println(car)

	car.color = "blue"

	fmt.Println("The car has been painted to the color: [", car.color, "]")

	//function
	showPlate(car.plate)

	//method
	car.displayPlate()

	fmt.Println("Choosing random plate...")
	car.randomPlate()

	car.displayPlate()
}

// function
func showPlate(p Plate) {
	fmt.Print("\n" + p.letters + "-" + p.numbers + "\n")
}

// method - struct Car
func (c Car) displayPlate() {
	fmt.Print("\n" + c.plate.letters + "-" + c.plate.numbers + "\n")
}

// It is necessary to use pointers to modify the actual struct rather than a copy of it.
func (c *Car) randomPlate() {
	c.plate = Plate{
		letters: c.plate.letters,
		numbers: strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)),
	}
}
