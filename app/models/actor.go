package models

import (
	"github.com/goravel/framework/support/carbon"
)

type Actor struct {
	ID         uint `gorm:"primaryKey;column:actor_id"`
	FirstName  string
	LastName   string
	LastUpdate carbon.DateTime
	Film       []*Film `gorm:"many2many:film_actor;joinForeignKey:actor_id;joinReferences:film_id"`
}

func (Actor) TableName() string {
	return "actor"
}
