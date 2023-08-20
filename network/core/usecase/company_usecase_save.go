package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type CompanyUseCaseSave struct {
	repository port.CompanyIRepository
}

func NewCompanyUseCaseSave(repository port.CompanyIRepository) *CompanyUseCaseSave {
	return &CompanyUseCaseSave{repository: repository}
}

func (o *CompanyUseCaseSave) Execute(Company *entity.Company) (*entity.Company, error) {

	if i, err := o.repository.Save(Company); err == nil && i != nil {
		return i.(*entity.Company), err
	} else {
		return nil, err
	}

}
