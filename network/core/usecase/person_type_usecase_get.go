package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type PersonTypeUseCaseGet struct {
	repository port.PersonTypeIRepository
}

func NewPersonTypeUseCaseGet(repository port.PersonTypeIRepository) *PersonTypeUseCaseGet {
	return &PersonTypeUseCaseGet{repository: repository}
}

func (o *PersonTypeUseCaseGet) Execute(id int64) (*entity.PersonType, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.PersonType), nil
	} else {
		return nil, err
	}
}
