package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/first_project/database"
	"github.com/prranavv/first_project/models"
)

type BookDTO struct {
	Book_ID int    `json:"book_id"`
	Author  string `json:"author"`
	ISBN    string `json:"isbn"`
}

func CreateBookResponse(book models.Book) BookDTO {
	return BookDTO{
		Book_ID: book.Book_ID,
		Author:  book.Author,
		ISBN:    book.ISBN,
	}
}

// POST
func Createbook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(err)
	}
	database.Database.Db.Create(&book)
	responsebook := CreateBookResponse(book)
	return c.Status(200).JSON(responsebook)
}

// GET
func Getbook(c *fiber.Ctx) error {
	books := []models.Book{}
	database.Database.Db.Find(&books)
	responsebooks := []BookDTO{}
	for _, book := range books {
		responsebook := CreateBookResponse(book)
		responsebooks = append(responsebooks, responsebook)
	}
	return c.Status(200).JSON(responsebooks)
}

func findbookbyid(book_id int, book *models.Book) error {
	database.Database.Db.Find(&book, "book_id=?", book_id)
	if book.Book_ID == 0 {
		return errors.New("book does not exist")
	}
	return nil
}

// UPDATE
func Updatebook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("book_id")
	if err != nil {
		c.Status(400).JSON("Please ensure that it is an integer")
	}
	var book models.Book
	if err := findbookbyid(id, &book); err != nil {
		return c.Status(400).SendString("over here")
	}
	type Updatebook struct {
		Author string `json:"author"`
		ISBN   string `json:"isbn"`
	}
	var newbook Updatebook
	if err := c.BodyParser(&newbook); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	book.Author = newbook.Author
	book.ISBN = newbook.ISBN
	database.Database.Db.Save(&book)
	responsebook := CreateBookResponse(book)
	return c.Status(200).JSON(responsebook)
}

// DELETE
func Deletebook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("book_id")
	if err != nil {
		return c.Status(400).JSON("Enter a number")
	}
	var book models.Book
	if err := findbookbyid(id, &book); err != nil {
		c.Status(400).JSON("Book does not exist")
	}
	if err := database.Database.Db.Delete(&book).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted")
}
