package entity

import (
	"manager/shared/types"
	"time"
)

type Regional struct {
	Id               int64           `json:"id"`
	IdParent         int64           `json:"id_parent"`
	Name             string          `json:"name"`
	ShortName        string          `json:"short_name"`
	Document         string          `json:"document"`
	DocumentType     string          `json:"document_type"`
	Telephone        string          `json:"telephone"`
	Email            string          `json:"email"`
	Address          types.JsonbType `json:"address"`
	Status           string          `json:"status"`
	Type             string          `json:"type"`
	Birthday         *time.Time      `json:"birthday"`
	CreationDateTime *time.Time      `json:"creation_data_time"`
	ChangeDateTime   *time.Time      `json:"change_data_time"`
}

func NewRegional() *Regional {
	return &Regional{
		Id:               0,
		IdParent:         0,
		Name:             "",
		ShortName:        "",
		Document:         "",
		DocumentType:     "",
		Telephone:        "",
		Email:            "",
		Address:          map[string]interface{}{},
		Status:           "",
		Type:             "",
		Birthday:         &time.Time{},
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (ent *Regional) GetId() int64 {
	return ent.Id
}

func (ent *Regional) SetId(id int64) {
	ent.Id = id
}

func (ent *Regional) IsValid() error {
	return nil
}
