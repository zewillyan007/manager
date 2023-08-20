package service

import "manager/network/core/domain/entity"

func FactoryCompany() *entity.Company {
	return entity.NewCompany()
}
