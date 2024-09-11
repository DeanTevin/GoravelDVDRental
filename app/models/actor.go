package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Actor struct {
	ID         uint `gorm:"primaryKey;column:actor_id"`
	FirstName  string
	LastName   string
	LastUpdate carbon.DateTime
}

func (Actor) TableName() string {
	return "actor"
}
