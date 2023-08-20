package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type CompanyUseCaseGetAll struct {
	repository port.CompanyIRepository
}

func NewCompanyUseCaseGetAll(repository port.CompanyIRepository) *CompanyUseCaseGetAll {
	return &CompanyUseCaseGetAll{repository: repository}
}

func (o *CompanyUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Company {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.Company, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.Company)
	}
	return entities
}
