package service

import (
	"fmt"
	"manager/network/core/domain/dto"
	"manager/network/core/port"
	"manager/network/core/usecase"
	port_shared "manager/shared/port"
	"manager/shared/types"
	"strconv"
)

type PersonTypeService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.PersonTypeIRepository
	ucGet      *usecase.PersonTypeUseCaseGet
	ucGetAll   *usecase.PersonTypeUseCaseGetAll
}

func NewPersonTypeService(provider port_shared.IResourceProvider) *PersonTypeService {
	repo := types.GetConstructor((*port.PersonTypeIRepository)(nil))(provider.GetDB()).(port.PersonTypeIRepository)
	repo.SetContext(provider.Context())

	return &PersonTypeService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewPersonTypeUseCaseGet(repo),
		ucGetAll:   usecase.NewPersonTypeUseCaseGetAll(repo),
	}
}

func (o *PersonTypeService) WithTransaction(transaction port_shared.ITransaction) *PersonTypeService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *PersonTypeService) Get(dtoIn *dto.PersonTypeDtoIn) (*dto.PersonTypeDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	PersonType, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewPersonTypeDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", PersonType.Id)
	dtoOut.Name = PersonType.Name
	dtoOut.Mnemonic = PersonType.Mnemonic
	dtoOut.Hint = PersonType.Hint

	if PersonType.CreationDateTime != nil {
		dtoOut.CreationDateTime = PersonType.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if PersonType.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = PersonType.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	if PersonType.DisableDateTime != nil {
		dtoOut.DisableDateTime = PersonType.DisableDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *PersonTypeService) GetAll(conditions ...interface{}) []*dto.PersonTypeDtoOut {

	var arrayPersonTypeDto []*dto.PersonTypeDtoOut

	arrayPersonType := o.ucGetAll.Execute(conditions...)

	for _, PersonType := range arrayPersonType {

		dtoOut := dto.NewPersonTypeDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", PersonType.Id)
		dtoOut.Name = PersonType.Name
		dtoOut.Mnemonic = PersonType.Mnemonic
		dtoOut.Hint = PersonType.Hint

		if PersonType.CreationDateTime != nil {
			dtoOut.CreationDateTime = PersonType.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if PersonType.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = PersonType.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		if PersonType.DisableDateTime != nil {
			dtoOut.DisableDateTime = PersonType.DisableDateTime.Format("2006-01-02 15:04:05")
		}

		arrayPersonTypeDto = append(arrayPersonTypeDto, dtoOut)
	}
	return arrayPersonTypeDto
}
