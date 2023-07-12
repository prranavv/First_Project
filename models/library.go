package models

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	Name      string `json:"name"`
	Address   string `json:"address"`
	BookRefer []int  `json:"book_id"`
	Books     []Book `gorm:"foreignKey:BookRefer"`
}
