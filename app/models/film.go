package models

import (
	"goravel/app/helpers"

	"github.com/goravel/framework/support/carbon"
)

type Film struct {
	ID              uint `gorm:"primaryKey;column:film_id"`
	Title           string
	Description     string
	ReleaseYear     int16
	LanguageID      int16
	RentalDuration  string
	RentalRate      float32
	Length          int16
	ReplacementCost float32
	Rating          string
	LastUpdate      carbon.DateTime
	SpecialFeatures helpers.SpecialFeatures `gorm:"type:text"` //using custom datatypes. This is great lol!
	Fulltext        helpers.Fulltext        `gorm:"type:text"`
	Language        *Language
}

func (Film) TableName() string {
	return "film"
}
