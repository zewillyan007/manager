package adapter

import (
	"manager/network/core/domain/entity"
	port2 "manager/network/core/port"
	"manager/shared/adapter"
	"manager/shared/port"
	"manager/shared/types"

	"gorm.io/gorm"
)

type RegionalRepository struct {
	adapter.RepositoryCRUD
}

func NewRegionalRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &RegionalRepository{}
	repo.EntityType = (*entity.Regional)(nil)
	repo.SetTable("network.regional")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.RegionalIRepository)(nil), func(args ...interface{}) interface{} {
		return NewRegionalRepository(args[0].(*gorm.DB))
	})
}
