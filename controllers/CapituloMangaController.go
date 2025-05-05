package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
)

func GetChapterURLHandler(c *fiber.Ctx) error {
	mangaURL := c.Query("url")
	cap := c.Query("cap")
	if mangaURL == "" || cap == "" {
		return c.Status(400).JSON(fiber.Map{
			"statusCode": 400,
			"data":       nil,
		})
	}

	scans, err := tumangaonline.FindChapterURLByNumber(mangaURL, cap)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"data":       nil,
		})
	}

	return c.JSON(models.ChapterMetadata{
		StatusCode: 200,
		Data:       scans,
	})
}
