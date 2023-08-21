package service

import (
	"encoding/json"
	"fmt"
	"manager/network/core/domain/dto"
	"manager/network/core/port"
	"manager/network/core/usecase"
	"manager/shared/grid"
	port_shared "manager/shared/port"
	"manager/shared/types"
	"strconv"
	"time"
)

type RegionalService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.RegionalIRepository
	ucGet      *usecase.RegionalUseCaseGet
	ucSave     *usecase.RegionalUseCaseSave
	ucGrid     *usecase.RegionalUseCaseGrid
	ucGetAll   *usecase.RegionalUseCaseGetAll
	ucRemove   *usecase.RegionalUseCaseRemove
}

func NewRegionalService(provider port_shared.IResourceProvider) *RegionalService {
	repo := types.GetConstructor((*port.RegionalIRepository)(nil))(provider.GetDB()).(port.RegionalIRepository)
	repo.SetContext(provider.Context())

	return &RegionalService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewRegionalUseCaseGet(repo),
		ucSave:     usecase.NewRegionalUseCaseSave(repo),
		ucGrid:     usecase.NewRegionalUseCaseGrid(repo),
		ucGetAll:   usecase.NewRegionalUseCaseGetAll(repo),
		ucRemove:   usecase.NewRegionalUseCaseRemove(repo),
	}
}

func (o *RegionalService) WithTransaction(transaction port_shared.ITransaction) *RegionalService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *RegionalService) Get(dtoIn *dto.RegionalDtoIn) (*dto.RegionalDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Regional, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewRegionalDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Regional.Id)
	dtoOut.IdParent = fmt.Sprintf("%d", Regional.IdParent)
	dtoOut.Name = Regional.Name
	dtoOut.ShortName = Regional.ShortName
	dtoOut.Document = Regional.Document
	dtoOut.DocumentType = Regional.DocumentType
	dtoOut.Telephone = Regional.Telephone
	dtoOut.Email = Regional.Email
	dtoOut.Status = Regional.Status
	dtoOut.Type = Regional.Type

	addressValue, err := Regional.Address.Value()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(addressValue.(string)), &dtoOut.Address)
	if err != nil {
		return nil, err
	}

	if Regional.Birthday != nil {
		dtoOut.Birthday = Regional.Birthday.Format("2006-01-02 15:04:05")
	}

	if Regional.CreationDateTime != nil {
		dtoOut.CreationDateTime = Regional.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if Regional.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Regional.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *RegionalService) GetAll(conditions ...interface{}) []*dto.RegionalDtoOut {

	var arrayRegionalDto []*dto.RegionalDtoOut

	arrayRegional := o.ucGetAll.Execute(conditions...)

	for _, Regional := range arrayRegional {

		dtoOut := dto.NewRegionalDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Regional.Id)
		dtoOut.IdParent = fmt.Sprintf("%d", Regional.IdParent)
		dtoOut.Name = Regional.Name
		dtoOut.ShortName = Regional.ShortName
		dtoOut.Document = Regional.Document
		dtoOut.DocumentType = Regional.DocumentType
		dtoOut.Telephone = Regional.Telephone
		dtoOut.Email = Regional.Email
		dtoOut.Status = Regional.Status
		dtoOut.Type = Regional.Type

		addressValue, _ := Regional.Address.Value()
		// if err != nil {
		// 	return nil, err
		// }

		_ = json.Unmarshal([]byte(addressValue.(string)), &dtoOut.Address)

		if Regional.Birthday != nil {
			dtoOut.Birthday = Regional.Birthday.Format("2006-01-02 15:04:05")
		}

		if Regional.CreationDateTime != nil {
			dtoOut.CreationDateTime = Regional.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if Regional.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Regional.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arrayRegionalDto = append(arrayRegionalDto, dtoOut)
	}
	return arrayRegionalDto
}

func (o *RegionalService) Save(dtoIn *dto.RegionalDtoIn) error {

	var err error
	Regional := FactoryRegional()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Regional.Id = int64(id)
	}

	Regional.IdParent, _ = strconv.ParseInt(dtoIn.IdParent, 10, 64)
	Regional.Name = dtoIn.Name
	Regional.ShortName = dtoIn.ShortName
	Regional.Document = dtoIn.Document
	Regional.DocumentType = dtoIn.DocumentType
	Regional.Telephone = dtoIn.Telephone
	Regional.Email = dtoIn.Email
	Regional.Status = dtoIn.Status
	Regional.Type = dtoIn.Type

	addressBytes, err := json.Marshal(dtoIn.Address)
	if err != nil {
		return err
	}
	Regional.Address.Scan(addressBytes)

	if len(dtoIn.Birthday) > 0 {
		birthday, err := time.Parse("2006-01-02 15:04:05", dtoIn.Birthday)
		if err != nil {
			return err
		}
		Regional.Birthday = &birthday
	}

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Regional.Id == 0 {
			Regional.CreationDateTime = &now
		} else {
			RegionalCurrent, _ := o.ucGet.Execute(Regional.Id)
			Regional.CreationDateTime = RegionalCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Regional.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Regional.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Regional.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(Regional)
	if err != nil {
		return err
	}
	return nil
}

func (o *RegionalService) Remove(dtoIn *dto.RegionalDtoIn) error {

	Regional := FactoryRegional()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Regional.Id = int64(id)
	}
	err := o.ucRemove.Execute(Regional)
	if err != nil {
		return err
	}

	return nil
}

func (o *RegionalService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
