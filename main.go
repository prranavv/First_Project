package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/first_project/database"
	"github.com/prranavv/first_project/routes"
)

func setuproutes(app *fiber.App) {
	app.Post("/book", routes.Createbook)
	app.Post("/lib", routes.CreateLibrary)
	app.Get("/book", routes.Getbooks)
	app.Get("/book/:book_id", routes.Getbook)
	app.Put("/book/:book_id", routes.Updatebook)
	app.Put("/lib", routes.InsertBookintoLibrary)
	app.Delete("/book/:book_id", routes.Deletebook)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setuproutes(app)
	log.Fatal(app.Listen(":3000"))
}
