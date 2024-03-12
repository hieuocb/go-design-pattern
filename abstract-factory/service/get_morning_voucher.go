package service

import (
	"abstract-factory/model"
)

type GetMorningVoucher struct{}

func (GetMorningVoucher) GetDrink() model.Drink {
	return model.Coffee{}
}

func (GetMorningVoucher) GetFood() model.Food {
	return model.Bread{}
}
