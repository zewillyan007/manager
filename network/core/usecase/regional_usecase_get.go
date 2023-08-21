package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type RegionalUseCaseGet struct {
	repository port.RegionalIRepository
}

func NewRegionalUseCaseGet(repository port.RegionalIRepository) *RegionalUseCaseGet {
	return &RegionalUseCaseGet{repository: repository}
}

func (o *RegionalUseCaseGet) Execute(id int64) (*entity.Regional, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.Regional), nil
	} else {
		return nil, err
	}
}
