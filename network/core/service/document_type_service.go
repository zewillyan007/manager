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

type DocumentTypeService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.DocumentTypeIRepository
	ucGet      *usecase.DocumentTypeUseCaseGet
	ucGetAll   *usecase.DocumentTypeUseCaseGetAll
}

func NewDocumentTypeService(provider port_shared.IResourceProvider) *DocumentTypeService {
	repo := types.GetConstructor((*port.DocumentTypeIRepository)(nil))(provider.GetDB()).(port.DocumentTypeIRepository)
	repo.SetContext(provider.Context())

	return &DocumentTypeService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewDocumentTypeUseCaseGet(repo),
		ucGetAll:   usecase.NewDocumentTypeUseCaseGetAll(repo),
	}
}

func (o *DocumentTypeService) WithTransaction(transaction port_shared.ITransaction) *DocumentTypeService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *DocumentTypeService) Get(dtoIn *dto.DocumentTypeDtoIn) (*dto.DocumentTypeDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	DocumentType, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewDocumentTypeDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", DocumentType.Id)
	dtoOut.Name = DocumentType.Name
	dtoOut.Mnemonic = DocumentType.Mnemonic
	dtoOut.Hint = DocumentType.Hint

	if DocumentType.CreationDateTime != nil {
		dtoOut.CreationDateTime = DocumentType.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if DocumentType.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DocumentType.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	if DocumentType.DisableDateTime != nil {
		dtoOut.DisableDateTime = DocumentType.DisableDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *DocumentTypeService) GetAll(conditions ...interface{}) []*dto.DocumentTypeDtoOut {

	var arrayDocumentTypeDto []*dto.DocumentTypeDtoOut

	arrayDocumentType := o.ucGetAll.Execute(conditions...)

	for _, DocumentType := range arrayDocumentType {

		dtoOut := dto.NewDocumentTypeDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", DocumentType.Id)
		dtoOut.Name = DocumentType.Name
		dtoOut.Mnemonic = DocumentType.Mnemonic
		dtoOut.Hint = DocumentType.Hint

		if DocumentType.CreationDateTime != nil {
			dtoOut.CreationDateTime = DocumentType.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if DocumentType.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DocumentType.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		if DocumentType.DisableDateTime != nil {
			dtoOut.DisableDateTime = DocumentType.DisableDateTime.Format("2006-01-02 15:04:05")
		}

		arrayDocumentTypeDto = append(arrayDocumentTypeDto, dtoOut)
	}
	return arrayDocumentTypeDto
}
