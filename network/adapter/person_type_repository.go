package adapter

import (
	"manager/network/core/domain/entity"
	port2 "manager/network/core/port"
	"manager/shared/adapter"
	"manager/shared/port"
	"manager/shared/types"

	"gorm.io/gorm"
)

type PersonTypeRepository struct {
	adapter.RepositoryCRUD
}

func NewPersonTypeRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &PersonTypeRepository{}
	repo.EntityType = (*entity.PersonType)(nil)
	repo.SetTable("domain.network_status")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.PersonTypeIRepository)(nil), func(args ...interface{}) interface{} {
		return NewPersonTypeRepository(args[0].(*gorm.DB))
	})
}
