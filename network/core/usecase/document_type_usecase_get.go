package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type DocumentTypeUseCaseGet struct {
	repository port.DocumentTypeIRepository
}

func NewDocumentTypeUseCaseGet(repository port.DocumentTypeIRepository) *DocumentTypeUseCaseGet {
	return &DocumentTypeUseCaseGet{repository: repository}
}

func (o *DocumentTypeUseCaseGet) Execute(id int64) (*entity.DocumentType, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.DocumentType), nil
	} else {
		return nil, err
	}
}
