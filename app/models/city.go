package models

import (
	"github.com/goravel/framework/support/carbon"
)

type City struct {
	ID         uint `gorm:"primaryKey;column:city_id"`
	City       string
	CountryID  int16
	LastUpdate carbon.DateTime
	Country    *Country
}

func (City) TableName() string {
	return "city"
}
