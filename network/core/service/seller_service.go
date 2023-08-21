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

type SellerService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.SellerIRepository
	ucGet      *usecase.SellerUseCaseGet
	ucSave     *usecase.SellerUseCaseSave
	ucGrid     *usecase.SellerUseCaseGrid
	ucGetAll   *usecase.SellerUseCaseGetAll
	ucRemove   *usecase.SellerUseCaseRemove
}

func NewSellerService(provider port_shared.IResourceProvider) *SellerService {
	repo := types.GetConstructor((*port.SellerIRepository)(nil))(provider.GetDB()).(port.SellerIRepository)
	repo.SetContext(provider.Context())

	return &SellerService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewSellerUseCaseGet(repo),
		ucSave:     usecase.NewSellerUseCaseSave(repo),
		ucGrid:     usecase.NewSellerUseCaseGrid(repo),
		ucGetAll:   usecase.NewSellerUseCaseGetAll(repo),
		ucRemove:   usecase.NewSellerUseCaseRemove(repo),
	}
}

func (o *SellerService) WithTransaction(transaction port_shared.ITransaction) *SellerService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *SellerService) Get(dtoIn *dto.SellerDtoIn) (*dto.SellerDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Seller, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewSellerDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Seller.Id)
	dtoOut.IdParent = fmt.Sprintf("%d", Seller.IdParent)
	dtoOut.Name = Seller.Name
	dtoOut.ShortName = Seller.ShortName
	dtoOut.Document = Seller.Document
	dtoOut.DocumentType = Seller.DocumentType
	dtoOut.Telephone = Seller.Telephone
	dtoOut.Email = Seller.Email
	dtoOut.Status = Seller.Status
	dtoOut.Type = Seller.Type

	addressValue, err := Seller.Address.Value()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(addressValue.(string)), &dtoOut.Address)
	if err != nil {
		return nil, err
	}

	if Seller.Birthday != nil {
		dtoOut.Birthday = Seller.Birthday.Format("2006-01-02 15:04:05")
	}

	if Seller.CreationDateTime != nil {
		dtoOut.CreationDateTime = Seller.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if Seller.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Seller.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *SellerService) GetAll(conditions ...interface{}) []*dto.SellerDtoOut {

	var arraySellerDto []*dto.SellerDtoOut

	arraySeller := o.ucGetAll.Execute(conditions...)

	for _, Seller := range arraySeller {

		dtoOut := dto.NewSellerDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Seller.Id)
		dtoOut.IdParent = fmt.Sprintf("%d", Seller.IdParent)
		dtoOut.Name = Seller.Name
		dtoOut.ShortName = Seller.ShortName
		dtoOut.Document = Seller.Document
		dtoOut.DocumentType = Seller.DocumentType
		dtoOut.Telephone = Seller.Telephone
		dtoOut.Email = Seller.Email
		dtoOut.Status = Seller.Status
		dtoOut.Type = Seller.Type

		addressValue, _ := Seller.Address.Value()
		// if err != nil {
		// 	return nil, err
		// }

		_ = json.Unmarshal([]byte(addressValue.(string)), &dtoOut.Address)

		if Seller.Birthday != nil {
			dtoOut.Birthday = Seller.Birthday.Format("2006-01-02 15:04:05")
		}

		if Seller.CreationDateTime != nil {
			dtoOut.CreationDateTime = Seller.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if Seller.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Seller.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arraySellerDto = append(arraySellerDto, dtoOut)
	}
	return arraySellerDto
}

func (o *SellerService) Save(dtoIn *dto.SellerDtoIn) error {

	var err error
	Seller := FactorySeller()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Seller.Id = int64(id)
	}

	Seller.IdParent, _ = strconv.ParseInt(dtoIn.IdParent, 10, 64)
	Seller.Name = dtoIn.Name
	Seller.ShortName = dtoIn.ShortName
	Seller.Document = dtoIn.Document
	Seller.DocumentType = dtoIn.DocumentType
	Seller.Telephone = dtoIn.Telephone
	Seller.Email = dtoIn.Email
	Seller.Status = dtoIn.Status
	Seller.Type = dtoIn.Type

	addressBytes, err := json.Marshal(dtoIn.Address)
	if err != nil {
		return err
	}
	Seller.Address.Scan(addressBytes)

	if len(dtoIn.Birthday) > 0 {
		birthday, err := time.Parse("2006-01-02 15:04:05", dtoIn.Birthday)
		if err != nil {
			return err
		}
		Seller.Birthday = &birthday
	}

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Seller.Id == 0 {
			Seller.CreationDateTime = &now
		} else {
			SellerCurrent, _ := o.ucGet.Execute(Seller.Id)
			Seller.CreationDateTime = SellerCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Seller.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Seller.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Seller.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(Seller)
	if err != nil {
		return err
	}
	return nil
}

func (o *SellerService) Remove(dtoIn *dto.SellerDtoIn) error {

	Seller := FactorySeller()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Seller.Id = int64(id)
	}
	err := o.ucRemove.Execute(Seller)
	if err != nil {
		return err
	}

	return nil
}

func (o *SellerService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
