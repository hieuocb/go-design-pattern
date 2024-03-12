package model

import "fmt"

type Coffee struct{}

func (Coffee) Drink() {
	fmt.Println("Drink Coffee")
}
