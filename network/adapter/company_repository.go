package adapter

import (
	"manager/network/core/domain/entity"
	port2 "manager/network/core/port"
	"manager/shared/adapter"
	"manager/shared/port"
	"manager/shared/types"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	adapter.RepositoryCRUD
}

func NewCompanyRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &CompanyRepository{}
	repo.EntityType = (*entity.Company)(nil)
	repo.SetTable("network.company")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.CompanyIRepository)(nil), func(args ...interface{}) interface{} {
		return NewCompanyRepository(args[0].(*gorm.DB))
	})
}
