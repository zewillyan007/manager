package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type SellerUseCaseGetAll struct {
	repository port.SellerIRepository
}

func NewSellerUseCaseGetAll(repository port.SellerIRepository) *SellerUseCaseGetAll {
	return &SellerUseCaseGetAll{repository: repository}
}

func (o *SellerUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Seller {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.Seller, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.Seller)
	}
	return entities
}
