package adapter

import (
	"manager/network/core/domain/entity"
	port2 "manager/network/core/port"
	"manager/shared/adapter"
	"manager/shared/port"
	"manager/shared/types"

	"gorm.io/gorm"
)

type SellerRepository struct {
	adapter.RepositoryCRUD
}

func NewSellerRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &SellerRepository{}
	repo.EntityType = (*entity.Seller)(nil)
	repo.SetTable("network.seller")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.SellerIRepository)(nil), func(args ...interface{}) interface{} {
		return NewSellerRepository(args[0].(*gorm.DB))
	})
}
