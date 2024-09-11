package models

import "github.com/goravel/framework/support/carbon"

type Rental struct {
	ID          uint `gorm:"primaryKey;column:rental_id"`
	RentalDate  carbon.DateTime
	InventoryID int
	CustomerID  int16
	ReturnDate  carbon.DateTime
	StaffID     int16
	LastUpdate  carbon.DateTime
	Amount      int
	Customer    *Customer
	Staff       *Staff
	Inventory   *Inventory
}

func (Rental) TableName() string {
	return "rental"
}
