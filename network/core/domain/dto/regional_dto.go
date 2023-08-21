package dto

type RegionalDtoIn struct {
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

func NewRegionalDtoIn() *RegionalDtoIn {

	return &RegionalDtoIn{
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

type RegionalDtoOut struct {
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

func NewRegionalDtoOut() *RegionalDtoOut {

	return &RegionalDtoOut{
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
