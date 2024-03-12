package service

import "abstract-factory/model"

type GetEveningVoucher struct{}

func (GetEveningVoucher) GetDrink() model.Drink {
	return model.Beer{}
}

func (GetEveningVoucher) GetFood() model.Food {
	return model.Octopus{}
}
