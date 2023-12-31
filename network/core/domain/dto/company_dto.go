package dto

type CompanyDtoIn struct {
	Id               string
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

func NewCompanyDtoIn() *CompanyDtoIn {

	return &CompanyDtoIn{
		Id:               "",
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

type CompanyDtoOut struct {
	Id               string
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

func NewCompanyDtoOut() *CompanyDtoOut {

	return &CompanyDtoOut{
		Id:               "",
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
