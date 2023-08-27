package models

import (
	"time"
)

type Search struct {
	SearchText string `json:"searchText"`
	Language   string `gorm:"column:language; varchar(100)" json:"language"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
}

type Movie struct {
	ID                   uint         `gorm:"primaryKey" json:"id"`
	MovieName            string       `gorm:"column:movie_name;type:varchar(100) ;not null" json:"movie_name"`
	Overview             string       `gorm:"column:overview";type:varchar(100) json:"overview"`
	Rating               int          `gorm:"column:rating;not null" json:"rating"`
	ReleaseDate          time.Time    `gorm:"column:release_date" json:"-"`
	Language             LanguageType `gorm:"column:language;type:ENUM('Hindi', 'English', 'Tamil', 'Telugu')" json:"language"`
	CreatedAt            time.Time    `gorm:"column:created_at" json:"-"`
	UpdatedAt            time.Time    `gorm:"column:updated_at" json:"-"`
	ReleaseDateFormatted string       `gorm:"" json:"release_date"`
}

// LanguageType represents the language enum type
type LanguageType string

const (
	Hindi   LanguageType = "Hindi"
	English LanguageType = "English"
	Tamil   LanguageType = "Tamil"
	Telugu  LanguageType = "Telugu"
)

// TableName specifies the table name for the Movie model
func (Movie) TableName() string {
	return "movie_db"
}
