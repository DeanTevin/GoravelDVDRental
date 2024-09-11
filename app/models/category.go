package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Category struct {
	ID         uint `gorm:"primaryKey;column:category_id"`
	Name       string
	LastUpdate carbon.DateTime
}

func (Category) TableName() string {
	return "category"
}
