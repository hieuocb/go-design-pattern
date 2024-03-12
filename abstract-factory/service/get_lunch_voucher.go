package service

import "abstract-factory/model"

type GetLunchVoucher struct{}

func (GetLunchVoucher) GetDrink() model.Drink {
	return model.MilkTea{}
}

func (GetLunchVoucher) GetFood() model.Food {
	return model.Cake{}
}
