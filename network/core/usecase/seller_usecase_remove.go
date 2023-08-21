package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type SellerUseCaseRemove struct {
	repository port.SellerIRepository
}

func NewSellerUseCaseRemove(repository port.SellerIRepository) *SellerUseCaseRemove {
	return &SellerUseCaseRemove{repository: repository}
}

func (o *SellerUseCaseRemove) Execute(Seller *entity.Seller) error {
	return o.repository.Remove(Seller)
}
