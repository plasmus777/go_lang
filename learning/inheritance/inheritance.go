package inheritance

import "fmt"

func Inheritance() {
	animal1 := Animal{
		Name:        "Rock dove",
		Sound:       "Coo",
		SoundNumber: 3,
	}

	pidgeon := Pidgeon{
		Animal: animal1,
		CanFly: true,
	}

	fmt.Println("--------", pidgeon.Name, "--------")
	pidgeon.MakeSounds()
	if pidgeon.CanFly {
		fmt.Println("Voosh voosh!")
	}

	//-------------------------------------------------

	animal2 := Animal{
		Name:        "Harrier",
		Sound:       "Bark",
		SoundNumber: 2,
	}

	dog := Dog{
		Animal: animal2,
		CanRun: true,
	}

	fmt.Println("--------", dog.Name, "--------")
	dog.MakeSounds()
	if dog.CanRun {
		fmt.Println("Arf arf!")
	}
}
