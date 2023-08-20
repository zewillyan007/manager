package dto

type CompanyDtoIn struct {
	Id               string
	Name             string
	ShortName        string
	Document         string
	Telephone        string
	Address          string
	CreationDateTime string
	ChangeDateTime   string
}

func NewCompanyDtoIn() *CompanyDtoIn {

	return &CompanyDtoIn{
		Id:               "",
		Name:             "",
		ShortName:        "",
		Document:         "",
		Telephone:        "",
		Address:          "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type CompanyDtoOut struct {
	Id               string
	Name             string
	ShortName        string
	Document         string
	Telephone        string
	Address          string
	CreationDateTime string
	ChangeDateTime   string
}

func NewCompanyDtoOut() *CompanyDtoOut {

	return &CompanyDtoOut{
		Id:               "",
		Name:             "",
		ShortName:        "",
		Document:         "",
		Telephone:        "",
		Address:          "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
