package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Inventory struct {
	ID         uint `gorm:"primaryKey;column:inventory_id"`
	FilmID     int16
	StoreID    int16
	LastUpdate carbon.DateTime
	Film       *Film
	Store      *Store
}

func (Inventory) TableName() string {
	return "inventory"
}
