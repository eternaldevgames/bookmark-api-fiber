package bookmark

import (
	"bookmark-api-fiber/database"

	"github.com/gofiber/fiber/v2"
)

func GetAllBookmarks(c *fiber.Ctx) error {
	result, err := database.GetAllBookmarks()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

func SaveBookmark(c *fiber.Ctx) error {
	newBookmark := new(database.Bookmark)

	err := c.BodyParser(newBookmark)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := database.CreateBookmark(newBookmark.Name, newBookmark.Url)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
