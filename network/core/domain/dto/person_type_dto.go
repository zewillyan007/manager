package dto

type PersonTypeDtoIn struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewPersonTypeDtoIn() *PersonTypeDtoIn {

	return &PersonTypeDtoIn{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}

type PersonTypeDtoOut struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewPersonTypeDtoOut() *PersonTypeDtoOut {

	return &PersonTypeDtoOut{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}
