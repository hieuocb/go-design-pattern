package model

import "fmt"

type MilkTea struct{}

func (MilkTea) Drink() {
	fmt.Println("Drink Milk Tea")
}
