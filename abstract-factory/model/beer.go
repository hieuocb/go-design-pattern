package model

import "fmt"

type Beer struct{}

func (Beer) Drink() {
	fmt.Println("Drink Beer!")
}
