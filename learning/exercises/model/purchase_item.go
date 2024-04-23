package model

import "fmt"

type Item struct {
	Name   string
	Number int
	Price  float64
}

func (i *Item) GetTotalValue() float64 {
	totalValue := float64(i.Number) * i.Price

	return totalValue
}

func (i *Item) PrintItem() {
	fmt.Println(i.Number, "X", i.Name, " - $", i.GetTotalValue(), " ($", i.Price, " each)")
}
