package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Category struct {
	ID         uint `gorm:"primaryKey;column:category_id"`
	Name       string
	LastUpdate carbon.DateTime
	Film       []*Film `gorm:"many2many:film_category;joinForeignKey:film_id;joinReferences:category_id"`
}

func (Category) TableName() string {
	return "category"
}
