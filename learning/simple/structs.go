package simple

import "fmt"

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
}

//function
func showPlate(p Plate) {
	fmt.Print("\n" + p.letters + "-" + p.numbers + "\n")
}

//method - struct Car
func (c Car) displayPlate() {
	fmt.Print("\n" + c.plate.letters + "-" + c.plate.numbers + "\n")
}
