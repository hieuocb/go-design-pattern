package service

import (
	"abstract-factory/model"
	"errors"
)

func GetVoucherFactory(campaignName string) (VoucherAbstractFactory, error) {
	switch campaignName {
	case "morning":
		return GetMorningVoucher{}, nil
	case "afternoon":
		return GetLunchVoucher{}, nil
	case "evening":
		return GetEveningVoucher{}, nil
	}

	return nil, errors.New("campaign not found")
}

func GetVoucher(factory VoucherAbstractFactory) model.Voucher {
	return model.Voucher{
		Drink: factory.GetDrink(),
		Food:  factory.GetFood(),
	}
}
