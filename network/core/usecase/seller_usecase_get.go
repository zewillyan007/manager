package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type SellerUseCaseGet struct {
	repository port.SellerIRepository
}

func NewSellerUseCaseGet(repository port.SellerIRepository) *SellerUseCaseGet {
	return &SellerUseCaseGet{repository: repository}
}

func (o *SellerUseCaseGet) Execute(id int64) (*entity.Seller, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.Seller), nil
	} else {
		return nil, err
	}
}
