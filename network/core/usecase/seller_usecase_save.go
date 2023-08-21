package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type SellerUseCaseSave struct {
	repository port.SellerIRepository
}

func NewSellerUseCaseSave(repository port.SellerIRepository) *SellerUseCaseSave {
	return &SellerUseCaseSave{repository: repository}
}

func (o *SellerUseCaseSave) Execute(Seller *entity.Seller) (*entity.Seller, error) {

	if i, err := o.repository.Save(Seller); err == nil && i != nil {
		return i.(*entity.Seller), err
	} else {
		return nil, err
	}
}
