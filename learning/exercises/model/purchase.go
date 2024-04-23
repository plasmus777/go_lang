package model

import (
	"errors"
	"fmt"
	"time"
)

type Purchase struct {
	Market       string
	PurchaseTime time.Time
	Items        map[string]Item
}

func InitializePurchase(name string) (*Purchase, error) {
	if len(name) == 0 {
		return nil, errors.New("The market's name cannot be empty.")
	}

	purchase := Purchase{
		Market:       name,
		PurchaseTime: time.Now(),
		Items:        make(map[string]Item),
	}

	return &purchase, nil
}

func (p *Purchase) AddItem1(i Item) {
	p.Items[i.Name] = i
}

func (p *Purchase) AddItem2(name string, number int, price float64) {
	p.Items[name] = Item{
		Name:   name,
		Number: number,
		Price:  price,
	}
}

func (p *Purchase) GetTotalValue() float64 {
	var totalValue float64

	for _, item := range p.Items {
		totalValue += item.GetTotalValue()
	}

	return totalValue
}

func (p *Purchase) ListItems() {
	for _, item := range p.Items {
		item.PrintItem()
	}
	fmt.Println("\nTotal: $", p.GetTotalValue())
}
