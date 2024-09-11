package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Language struct {
	ID         uint `gorm:"primaryKey;column:language_id"`
	Name       string
	LastUpdate carbon.DateTime
}

func (Language) TableName() string {
	return "language"
}
