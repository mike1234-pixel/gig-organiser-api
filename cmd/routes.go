package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mike1234-pixel/gig-organiser-api/handlers"
)

func setupRoutes(app *fiber.App) {
	// Use the CORS middleware to enable CORS for localhost:3000 and localhost:3001
	app.Use(cors.New(cors.Config{
		// Use the correct syntax to assign a value to the AllowOrigins field
		AllowOrigins:     "http://localhost:3000, http://localhost:3001",
		AllowCredentials: true,
	}))

	app.Get("/jobs", handlers.ListJobs)

	app.Post("/job", handlers.CreateJob)

	app.Delete("/jobs/:id", handlers.DeleteJob)

	app.Put("/jobs/:id", handlers.UpdateJob)
}
