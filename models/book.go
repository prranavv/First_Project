package models

import "time"

type Book struct {
	Book_ID   int `json:"book_id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"name"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
}
