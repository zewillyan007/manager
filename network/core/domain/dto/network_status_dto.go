package dto

type NetworkStatusDtoIn struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewNetworkStatusDtoIn() *NetworkStatusDtoIn {

	return &NetworkStatusDtoIn{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}

type NetworkStatusDtoOut struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewNetworkStatusDtoOut() *NetworkStatusDtoOut {

	return &NetworkStatusDtoOut{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}
