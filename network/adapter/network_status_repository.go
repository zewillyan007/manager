package adapter

import (
	"manager/network/core/domain/entity"
	port2 "manager/network/core/port"
	"manager/shared/adapter"
	"manager/shared/port"
	"manager/shared/types"

	"gorm.io/gorm"
)

type NetworkStatusRepository struct {
	adapter.RepositoryCRUD
}

func NewNetworkStatusRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &NetworkStatusRepository{}
	repo.EntityType = (*entity.NetworkStatus)(nil)
	repo.SetTable("domain.network_status")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.NetworkStatusIRepository)(nil), func(args ...interface{}) interface{} {
		return NewNetworkStatusRepository(args[0].(*gorm.DB))
	})
}
