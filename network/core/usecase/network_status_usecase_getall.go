package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type NetworkStatusUseCaseGetAll struct {
	repository port.NetworkStatusIRepository
}

func NewNetworkStatusUseCaseGetAll(repository port.NetworkStatusIRepository) *NetworkStatusUseCaseGetAll {
	return &NetworkStatusUseCaseGetAll{repository: repository}
}

func (o *NetworkStatusUseCaseGetAll) Execute(conditions ...interface{}) []*entity.NetworkStatus {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.NetworkStatus, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.NetworkStatus)
	}
	return entities
}
