package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	City string `json:"city,omitempty"`

	PinCode string `json:"pinCode,omitempty"`

	Street string `json:"street,omitempty"`
}
