package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/procode2/structio/handlers"
)

func SetupRoutes(app *fiber.App) {
	home := app.Group("/")
	home.Get("/", handlers.HandleGetHome)

	path := app.Group("/paths")
	path.Get("/", handlers.HandleGetExplorePaths)
	path.Post("/", handlers.HandlePostPath)
	path.Get("/:pathId", handlers.HandleGetPath)
	path.Put("/:pathid", handlers.HandleUpdatePath)

	auth := app.Group("/auth")
	auth.Get("/signin", handlers.HandleGetSingin)
	auth.Post("/signin", handlers.HandlePostSingin)
	auth.Get("/login", handlers.HandleGetLogin)
	auth.Post("/login", handlers.HandlePostLogin)
	auth.Delete("/logout", handlers.HandleDeleteLogout)

}
