package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type CompanyUseCaseGet struct {
	repository port.CompanyIRepository
}

func NewCompanyUseCaseGet(repository port.CompanyIRepository) *CompanyUseCaseGet {
	return &CompanyUseCaseGet{repository: repository}
}

func (o *CompanyUseCaseGet) Execute(id int64) (*entity.Company, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.Company), nil
	} else {
		return nil, err
	}
}

//oie
