package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Customer struct {
	ID         uint `gorm:"primaryKey;column:customer_id"`
	StoreID    int16
	FirstName  string
	LastName   string
	Email      string
	AddressID  int16
	Activebool bool
	CreateDate carbon.Date
	LastUpdate carbon.DateTime
	Active     int16
	Address    *Address
	Store      *Store
}

func (Customer) TableName() string {
	return "customer"
}
