package service

import "manager/network/core/domain/entity"

func FactoryCompany() *entity.Company {
	return entity.NewCompany()
}

func FactoryRegional() *entity.Regional {
	return entity.NewRegional()
}

func FactorySeller() *entity.Seller {
	return entity.NewSeller()
}

func FactoryNetworkStatus() *entity.NetworkStatus {
	return entity.NewNetworkStatus()
}

func FactoryDocumentType() *entity.DocumentType {
	return entity.NewDocumentType()
}
