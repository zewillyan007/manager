package adapter

import (
	"manager/network/core/domain/entity"
	port2 "manager/network/core/port"
	"manager/shared/adapter"
	"manager/shared/port"
	"manager/shared/types"

	"gorm.io/gorm"
)

type DocumentTypeRepository struct {
	adapter.RepositoryCRUD
}

func NewDocumentTypeRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &DocumentTypeRepository{}
	repo.EntityType = (*entity.DocumentType)(nil)
	repo.SetTable("domain.document_type")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.DocumentTypeIRepository)(nil), func(args ...interface{}) interface{} {
		return NewDocumentTypeRepository(args[0].(*gorm.DB))
	})
}
