package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Address struct {
	ID              uint `gorm:"primaryKey;column:address_id"`
	Address         string
	AddressOptional string `gorm:"column:address2"`
	District        string
	CityID          int16
	PostalCode      string
	Phone           string
	LastUpdate      carbon.DateTime
}

func (Address) TableName() string {
	return "address"
}
