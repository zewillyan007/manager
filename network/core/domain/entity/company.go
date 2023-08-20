package entity

import (
	"manager/shared/types"
	"time"
)

type Company struct {
	Id               int64           `json:"id"`
	Name             string          `json:"name"`
	ShortName        string          `json:"short_name"`
	Document         string          `json:"document"`
	Telephone        string          `json:"telephone"`
	Address          types.JsonbType `json:"address"`
	CreationDateTime *time.Time      `json:"creation_data_time"`
	ChangeDateTime   *time.Time      `json:"change_data_time"`
	// Ddd               string     `json:"ddd"`
	// Email             string     `json:"email"`
}

func NewCompany() *Company {
	return &Company{
		Id:               0,
		Name:             "",
		ShortName:        "",
		Document:         "",
		Telephone:        "",
		Address:          map[string]interface{}{},
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (ent *Company) GetId() int64 {
	return ent.Id
}

func (ent *Company) SetId(id int64) {
	ent.Id = id
}

func (ent *Company) IsValid() error {
	return nil
}
