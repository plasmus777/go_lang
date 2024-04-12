package fluxcontrol

import "fmt"

func If() {

	value := 999.00
	var valuePlus float64

	valuePlus = value

	/*
		<
		>
		<=
		>=
	*/

	addExtra := true

	if addExtra {
		fmt.Println("Adding 100 to the total value.")
		if value <= 1000 {
			valuePlus += 100
		}
	} else {
		fmt.Println("The total value will remain intact.")
	}

	if valuePlus > 1000 {
		fmt.Println("The total value is greater than 1000.")
	} else if valuePlus == 1000 {
		fmt.Println("The total value is equal to 1000.")
	} else if valuePlus < 1000 {
		fmt.Println("The total value is lesser than 1000.")
	}

	fmt.Println(valuePlus)
}
