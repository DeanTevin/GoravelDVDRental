package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Staff struct {
	ID         uint `gorm:"primaryKey;column:staff_id"`
	FirstName  string
	LastName   string
	AddressID  int16
	Email      string
	StoreID    int16
	Active     bool `json:"-"`
	Username   string
	Password   string `json:"-"`
	LastUpdate carbon.DateTime
	Picture    []byte
	Store      *Store   `gorm:"ForeignKey:ID"`
	Address    *Address `gorm:"ForeignKey:ID"`
}

func (Staff) TableName() string {
	return "staff"
}
