package service

import (
	"fmt"
	"manager/network/core/domain/dto"
	"manager/network/core/port"
	"manager/network/core/usecase"
	"manager/shared/grid"
	"manager/shared/helper"
	port_shared "manager/shared/port"
	"manager/shared/types"
	"strconv"
	"time"
)

type CompanyService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.CompanyIRepository
	ucGet      *usecase.CompanyUseCaseGet
	ucSave     *usecase.CompanyUseCaseSave
	ucGrid     *usecase.CompanyUseCaseGrid
	ucGetAll   *usecase.CompanyUseCaseGetAll
	ucRemove   *usecase.CompanyUseCaseRemove
}

func NewCompanyService(provider port_shared.IResourceProvider) *CompanyService {
	repo := types.GetConstructor((*port.CompanyIRepository)(nil))(provider.GetDB()).(port.CompanyIRepository)
	repo.SetContext(provider.Context())

	return &CompanyService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewCompanyUseCaseGet(repo),
		ucSave:     usecase.NewCompanyUseCaseSave(repo),
		ucGrid:     usecase.NewCompanyUseCaseGrid(repo),
		ucGetAll:   usecase.NewCompanyUseCaseGetAll(repo),
		ucRemove:   usecase.NewCompanyUseCaseRemove(repo),
	}
}

func (o *CompanyService) WithTransaction(transaction port_shared.ITransaction) *CompanyService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *CompanyService) Get(dtoIn *dto.CompanyDtoIn) (*dto.CompanyDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Company, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewCompanyDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Company.Id)
	dtoOut.Name = Company.Name
	dtoOut.ShortName = Company.ShortName
	dtoOut.Document = Company.Document
	dtoOut.Telephone = Company.Telephone
	dtoOut.Address = Company.Address

	if Company.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("Y-m-d H:i:s", *Company.CreationDateTime)
	}

	if Company.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("Y-m-d H:i:s", *Company.ChangeDateTime)
	}

	return dtoOut, nil
}

func (o *CompanyService) GetAll(conditions ...interface{}) []*dto.CompanyDtoOut {

	var arrayCompanyDto []*dto.CompanyDtoOut

	arrayCompany := o.ucGetAll.Execute(conditions...)

	for _, Company := range arrayCompany {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewCompanyDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Company.Id)
		dtoOut.Name = Company.Name
		dtoOut.ShortName = Company.ShortName
		dtoOut.Document = Company.Document
		dtoOut.Telephone = Company.Telephone
		dtoOut.Address = Company.Address

		if Company.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("Y-m-d H:i:s", *Company.CreationDateTime)
		}

		if Company.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("Y-m-d H:i:s", *Company.ChangeDateTime)
		}

		arrayCompanyDto = append(arrayCompanyDto, dtoOut)
	}
	return arrayCompanyDto
}

func (o *CompanyService) Save(dtoIn *dto.CompanyDtoIn) error {

	Company := FactoryCompany()
	DateHelper := helper.NewDateHelper()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Company.Id = int64(id)
	}

	Company.Name = dtoIn.Name
	Company.ShortName = dtoIn.ShortName
	Company.Document = dtoIn.Document
	Company.Telephone = dtoIn.Telephone
	Company.Address = dtoIn.Address

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Company.Id == 0 {
			Company.CreationDateTime = &now
		} else {
			CompanyCurrent, _ := o.ucGet.Execute(Company.Id)
			Company.CreationDateTime = CompanyCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := DateHelper.Parse("Y-m-d H:i:s", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Company.CreationDateTime = CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Company.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := DateHelper.Parse("Y-m-d H:i:s", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Company.ChangeDateTime = ChangeDateTime
	}

	_, err := o.ucSave.Execute(Company)
	if err != nil {
		return err
	}
	return nil
}

func (o *CompanyService) Remove(dtoIn *dto.CompanyDtoIn) error {

	Company := FactoryCompany()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Company.Id = int64(id)
	}
	err := o.ucRemove.Execute(Company)
	if err != nil {
		return err
	}

	return nil
}

func (o *CompanyService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
