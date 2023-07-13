package models

import "time"

type Library struct {
	Lib_ID    uint `json:"lib_id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"name"`
	Address   string `json:"address"`
	BookRefer int    `json:"book_id"`
	Books     []Book `gorm:"foreignKey:BookRefer"`
}
