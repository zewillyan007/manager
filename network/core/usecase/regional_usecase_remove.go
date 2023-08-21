package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type RegionalUseCaseRemove struct {
	repository port.RegionalIRepository
}

func NewRegionalUseCaseRemove(repository port.RegionalIRepository) *RegionalUseCaseRemove {
	return &RegionalUseCaseRemove{repository: repository}
}

func (o *RegionalUseCaseRemove) Execute(Regional *entity.Regional) error {
	return o.repository.Remove(Regional)
}
