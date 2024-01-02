package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Address Address `gorm:"foreignKey:ID" json:"address,omitempty"`

	Age int `json:"age,omitempty"`

	Name string `json:"name,omitempty"`

	Sign rune `json:"sign,omitempty"`
}
