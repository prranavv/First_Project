package models

type Library struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	BookRefer []int  `json:"book_id"`
	Books     []Book `gorm:"foreignKey:BookRefer"`
}
