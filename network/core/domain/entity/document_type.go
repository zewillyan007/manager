package entity

import (
	"time"
)

type DocumentType struct {
	Id               int64      `json:"id"`
	Name             string     `json:"name"`
	Mnemonic         string     `json:"mnemonic"`
	Hint             string     `json:"hint"`
	CreationDateTime *time.Time `json:"creation_data_time"`
	ChangeDateTime   *time.Time `json:"change_data_time"`
	DisableDateTime  *time.Time `json:"disable_data_time"`
}

func NewDocumentType() *DocumentType {
	return &DocumentType{
		Id:               0,
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
		DisableDateTime:  &time.Time{},
	}
}

func (ent *DocumentType) GetId() int64 {
	return ent.Id
}

func (ent *DocumentType) SetId(id int64) {
	ent.Id = id
}

func (ent *DocumentType) IsValid() error {
	return nil
}
