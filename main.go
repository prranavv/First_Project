package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/first_project/database"
	"github.com/prranavv/first_project/routes"
)

func setuproutes(app *fiber.App) {
	app.Post("/book", routes.Createbook)
	app.Get("/book", routes.Getbook)
	app.Put("/book/:book_id", routes.Updatebook)
	app.Delete("/book/:book_id", routes.Deletebook)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setuproutes(app)
	log.Fatal(app.Listen(":3000"))
}
