package dto

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
