package usecase

import (
	"manager/network/core/domain/entity"
	"manager/network/core/port"
)

type DocumentTypeUseCaseGetAll struct {
	repository port.DocumentTypeIRepository
}

func NewDocumentTypeUseCaseGetAll(repository port.DocumentTypeIRepository) *DocumentTypeUseCaseGetAll {
	return &DocumentTypeUseCaseGetAll{repository: repository}
}

func (o *DocumentTypeUseCaseGetAll) Execute(conditions ...interface{}) []*entity.DocumentType {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.DocumentType, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.DocumentType)
	}
	return entities
}
