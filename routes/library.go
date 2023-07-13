package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/first_project/database"
	"github.com/prranavv/first_project/models"
)

type LibraryDTO struct {
	Lib_ID  uint      `json:"lib_id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Books   []BookDTO `json:"books"`
}

func CreateLibResponse(lib models.Library, books []BookDTO) LibraryDTO {
	return LibraryDTO{
		Lib_ID:  lib.Lib_ID,
		Name:    lib.Name,
		Address: lib.Address,
		Books:   books,
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

func findlibbyid(id int, lib *models.Library) error {
	database.Database.Db.Find(&lib, "id=?", id)
	if lib.Lib_ID == 0 {
		return errors.New("library does not exist")
	}
	return nil
}

type InsertBook struct {
	Book_id int `query:"book_id"`
	Lib_id  int `query:"lib_id"`
}

func InsertBookintoLibrary(c *fiber.Ctx) error {
	i := new(InsertBook)
	if err := c.QueryParser(i); err != nil {
		return c.SendString("Cant parse query!")
	}
	var book models.Book
	if err := findbookbyid(i.Book_id, &book); err != nil {
		return c.Status(400).JSON("Cant find book")
	}
	var lib models.Library
	if err := findlibbyid(i.Lib_id, &lib); err != nil {
		return c.Status(400).SendString("Cant find lib")
	}
	//lib.Books = append(lib.Books, book)
	database.Database.Db.Save(&lib)
	var responsebooks []BookDTO
	responsebook := CreateBookResponse(book)
	responsebooks = append(responsebooks, responsebook)
	responselib := CreateLibResponse(lib, responsebooks)
	return c.JSON(responselib)
}

// lib/:lid_id/:book_id
// func InsertBookintoLibrary(c *fiber.Ctx) error {
// 	m := c.Queries()
// 	var book models.Book
// 	if err := c.BodyParser(&book); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	intvar, err := strconv.Atoi(m["book_id"])
// 	if err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	if err := findbookbyid(intvar, &book); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	var lib models.Library
// 	intvar3, err := strconv.Atoi(m["lib_id"])
// 	if err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	if err := findlibbyid(intvar3, &lib); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	lib.Books = append(lib.Books, book)
// 	database.Database.Db.Save(&lib)
// 	var responsebooks []BookDTO
// 	responsebooks = CreateBookResponse(book)
// 	responselib := CreateLibResponse(lib, responsebook)
// 	return c.Status(200).JSON(responselib)
// }
