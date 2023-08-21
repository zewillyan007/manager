package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type NetworkStatusUseCaseGet struct {
	repository port.NetworkStatusIRepository
}

func NewNetworkStatusUseCaseGet(repository port.NetworkStatusIRepository) *NetworkStatusUseCaseGet {
	return &NetworkStatusUseCaseGet{repository: repository}
}

func (o *NetworkStatusUseCaseGet) Execute(id int64) (*entity.NetworkStatus, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.NetworkStatus), nil
	} else {
		return nil, err
	}
}
