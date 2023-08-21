package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type PersonTypeUseCaseGetAll struct {
	repository port.PersonTypeIRepository
}

func NewPersonTypeUseCaseGetAll(repository port.PersonTypeIRepository) *PersonTypeUseCaseGetAll {
	return &PersonTypeUseCaseGetAll{repository: repository}
}

func (o *PersonTypeUseCaseGetAll) Execute(conditions ...interface{}) []*entity.PersonType {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.PersonType, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.PersonType)
	}
	return entities
}
