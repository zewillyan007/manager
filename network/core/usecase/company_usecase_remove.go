package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type CompanyUseCaseRemove struct {
	repository port.CompanyIRepository
}

func NewCompanyUseCaseRemove(repository port.CompanyIRepository) *CompanyUseCaseRemove {
	return &CompanyUseCaseRemove{repository: repository}
}

func (o *CompanyUseCaseRemove) Execute(Company *entity.Company) error {
	return o.repository.Remove(Company)
}
