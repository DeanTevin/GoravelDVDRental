package models

import "github.com/goravel/framework/support/carbon"

type Payment struct {
	ID         uint `gorm:"primaryKey;column:payment_id"`
	CustomerID int16
	StaffID    int16
	RentalID   int16
	LastUpdate carbon.DateTime
	Amount     int
	Customer   *Customer
	Staff      *Staff
}

func (Payment) TableName() string {
	return "payment"
}
