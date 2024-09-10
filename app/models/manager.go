package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Manager struct {
	ID         uint `gorm:"primaryKey;column:staff_id"`
	FirstName  string
	LastName   string
	AddressId  int16
	Email      string
	Active     bool `json:"-"`
	Username   string
	Password   string `json:"-"`
	LastUpdate carbon.DateTime
	Picture    []byte
}

func (Manager) TableName() string {
	return "staff"
}
