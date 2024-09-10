package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Store struct {
	ID         uint    `gorm:"primaryKey;column:store_id"`
	Manager    Manager `gorm:"ForeignKey:ID"`
	AddressId  int16
	LastUpdate carbon.DateTime
	Staff      []*Staff
}

func (Store) TableName() string {
	return "store"
}
