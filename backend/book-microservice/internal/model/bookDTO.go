package model

type BookDTO struct {
	Name        string `gorm:"not null;"`
	Author      string `gorm:"not null;"`
	Description string `gorm:"not null;"`
	Image       string `gorm:"not null;"`
}
