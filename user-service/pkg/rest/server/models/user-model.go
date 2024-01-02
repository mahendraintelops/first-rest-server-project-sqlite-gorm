package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	Address Address `gorm:"foreignKey:ID" json:"address,omitempty"`

	Age int `json:"age,omitempty"`

	Name string `json:"name,omitempty"`

	Sign rune `json:"sign,omitempty"`
}
