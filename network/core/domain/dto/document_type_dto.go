package dto

type DocumentTypeDtoIn struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewDocumentTypeDtoIn() *DocumentTypeDtoIn {

	return &DocumentTypeDtoIn{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}

type DocumentTypeDtoOut struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewDocumentTypeDtoOut() *DocumentTypeDtoOut {

	return &DocumentTypeDtoOut{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}
