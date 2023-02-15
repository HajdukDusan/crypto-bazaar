package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Name        string `gorm:"not null;"`
	Author      string `gorm:"not null;"`
	Description string `gorm:"not null;"`
	Image       string `gorm:"not null;"`
}

func (Book) TableName() string {
	return "Books"
}
