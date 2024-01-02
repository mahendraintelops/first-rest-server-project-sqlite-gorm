package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primaryKey;primaryKey"`

	Address Address `gorm:"foreignKey:ID" json:"address,omitempty"`

	Age int `json:"age,omitempty"`

	Name string `json:"name,omitempty"`

	Sign rune `json:"sign,omitempty"`
}

// BeforeCreate hook to generate UUID before creating a record
func (m *User) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	return nil
}
