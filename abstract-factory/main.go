package main

import (
	"abstract-factory/service"
	"log"
)

func main() {
	voucherFactory, err := service.GetVoucherFactory("morning")

	if err != nil {
		log.Fatal(err)
	}

	myVoucher := service.GetVoucher(voucherFactory)

	myVoucher.Drink.Drink()
	myVoucher.Food.Eat()

}
