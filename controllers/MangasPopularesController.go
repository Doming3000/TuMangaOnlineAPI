package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
)

func GetMangasPopularesWithPagination(c *fiber.Ctx) error {

	param := c.Query("pageNumber", "0")
	log.Println("Página solicitada:", param)

	id, err := strconv.Atoi(param)
	if err != nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "error pageNumber no puede ser null",
		}
		return c.JSON(response)
	}

	var mangasPopulares []models.Manga
	if id > 0 {
		log.Printf("Obteniendo mangas populares de la página %d", id)
		mangasPopulares = tumangaonline.GetMangasPopulares(id)
	} else {
		log.Println("Página no válida o 0, obteniendo mangas populares de la página 1")
		mangasPopulares = tumangaonline.GetMangasPopulares(1)
	}

	if len(mangasPopulares) == 0 {
		log.Println("No se encontraron mangas populares.")
		response := models.Response{
			StatusCode: http.StatusNotFound,
			Data:       "No se encontraron mangas populares en esta página.",
		}
		return c.JSON(response)
	}

	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       mangasPopulares,
	}
	return c.JSON(response)
}
