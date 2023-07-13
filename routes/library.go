package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/first_project/database"
	"github.com/prranavv/first_project/models"
)

type LibraryDTO struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Books   []BookDTO `json:"books"`
}

func CreateLibResponse(lib models.Library, book []BookDTO) LibraryDTO {
	return LibraryDTO{
		ID:      lib.ID,
		Name:    lib.Name,
		Address: lib.Address,
		Books:   book,
	}
}

func CreateLibrary(c *fiber.Ctx) error {
	var lib models.Library
	if err := c.BodyParser(&lib); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&lib)
	var book []BookDTO
	responselib := CreateLibResponse(lib, book)
	return c.Status(200).JSON(responselib)
}
