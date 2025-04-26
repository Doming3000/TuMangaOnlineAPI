package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/controllers"
)

func main() {
	startServer()
}

func startServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		paths := []string{
			"/api/v1/manga/populares",
			"/api/v1/manga/populares-josei",
			"/api/v1/manga/populares-seinen",
			"/api/v1/manga/info",
			"/api/v1/manga/library",
			"/api/v1/manga/listas",
			"/api/v1/get-cookies",
			"/api/v1/get-manga",
		}
		return c.JSON(paths)
	})

	app.Get("/api/v1/manga/populares", controllers.GetMangasPopularesWithPagination)
	app.Get("/api/v1/manga/populares-josei", controllers.GetMangasPopularesJosei)
	app.Get("/api/v1/manga/populares-seinen", controllers.GetMangasPopularesSeinen)
	app.Get("/api/v1/manga/info", controllers.GetInfoManga)
	app.Get("/api/v1/manga/library", controllers.GetInfoLibrary)
	app.Get("/api/v1/manga/listas", controllers.GetListasMangas)
	app.Get("/api/v1/get-cookies", controllers.GetCookiesFromTMO)
	app.Get("/api/v1/get-manga", controllers.GetPageFromTMOWithCookie)

	port := getPort()
	addr := "0.0.0.0:" + port
	log.Printf("Starting server on %s\n", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return port
}
