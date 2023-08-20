package dto

type CompanyDtoIn struct {
	Id               string
	Name             string
	ShortName        string
	Document         string
	Telephone        string
	Address          *Address
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
		Address:          &Address{},
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
	Address          *Address
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
		Address:          &Address{},
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type Address struct {
	Street  string
	Number  string
	City    string
	State   string
	Zip     string
	Country string
}

func NewAddress() *Address {
	return &Address{
		Street:  "",
		Number:  "",
		City:    "",
		State:   "",
		Zip:     "",
		Country: "",
	}
}
