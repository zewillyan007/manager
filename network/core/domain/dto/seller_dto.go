package dto

type SellerDtoIn struct {
	Id               string
	IdParent         string
	Name             string
	ShortName        string
	Document         string
	DocumentType     string
	Telephone        string
	Email            string
	Address          *Address
	Status           string
	Type             string
	Birthday         string
	CreationDateTime string
	ChangeDateTime   string
}

func NewSellerDtoIn() *SellerDtoIn {

	return &SellerDtoIn{
		Id:               "",
		IdParent:         "",
		Name:             "",
		ShortName:        "",
		Document:         "",
		DocumentType:     "",
		Telephone:        "",
		Email:            "",
		Address:          &Address{},
		Status:           "",
		Type:             "",
		Birthday:         "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type SellerDtoOut struct {
	Id               string
	IdParent         string
	Name             string
	ShortName        string
	Document         string
	DocumentType     string
	Telephone        string
	Email            string
	Address          *Address
	Status           string
	Type             string
	Birthday         string
	CreationDateTime string
	ChangeDateTime   string
}

func NewSellerDtoOut() *SellerDtoOut {

	return &SellerDtoOut{
		Id:               "",
		IdParent:         "",
		Name:             "",
		ShortName:        "",
		Document:         "",
		DocumentType:     "",
		Telephone:        "",
		Email:            "",
		Address:          &Address{},
		Status:           "",
		Type:             "",
		Birthday:         "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
