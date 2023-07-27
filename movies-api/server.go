package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Marbex4x4/lab-ic-ms-2023/movies-api/database"	
	"github.com/Marbex4x4/lab-ic-ms-2023/movies-api/handlers"
)

func server(bind, port string) {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(bind + ":" + port)
}

func setupRoutes(app *fiber.App) {
	app.Get("/movies", handlers.ListMovies)
	app.Get("/movies/:id", handlers.GetMovie)
}
