package inheritance

import "fmt"

type Animal struct {
	Name        string
	Sound       string
	SoundNumber int
}

func (a *Animal) MakeSounds() {
	for i := 0; i < a.SoundNumber; i++ {
		fmt.Println(a.Sound)
	}
}

type Pidgeon struct {
	Animal
	CanFly bool
}

type Dog struct {
	Animal
	CanRun bool
}
