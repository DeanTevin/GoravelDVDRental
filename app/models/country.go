package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Country struct {
	ID         uint `gorm:"primaryKey;column:country_id"`
	Country    string
	LastUpdate carbon.DateTime
}

func (Country) TableName() string {
	return "country"
}
