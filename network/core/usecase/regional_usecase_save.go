package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type RegionalUseCaseSave struct {
	repository port.RegionalIRepository
}

func NewRegionalUseCaseSave(repository port.RegionalIRepository) *RegionalUseCaseSave {
	return &RegionalUseCaseSave{repository: repository}
}

func (o *RegionalUseCaseSave) Execute(Regional *entity.Regional) (*entity.Regional, error) {

	if i, err := o.repository.Save(Regional); err == nil && i != nil {
		return i.(*entity.Regional), err
	} else {
		return nil, err
	}
}
