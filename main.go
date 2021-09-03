package main

import (
	"bookmark-api-fiber/bookmark"
	"bookmark-api-fiber/database"

	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/bookmark", bookmark.GetAllBookmarks)
	app.Post("/api/bookmark", bookmark.SaveBookmark)
}

func main() {
	app := fiber.New()
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
