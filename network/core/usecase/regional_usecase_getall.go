package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type RegionalUseCaseGetAll struct {
	repository port.RegionalIRepository
}

func NewRegionalUseCaseGetAll(repository port.RegionalIRepository) *RegionalUseCaseGetAll {
	return &RegionalUseCaseGetAll{repository: repository}
}

func (o *RegionalUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Regional {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.Regional, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.Regional)
	}
	return entities
}
