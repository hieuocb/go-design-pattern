package service

import "abstract-factory/model"

type VoucherAbstractFactory interface {
	GetDrink() model.Drink
	GetFood() model.Food
}
