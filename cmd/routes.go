package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mike1234-pixel/gig-organiser-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/jobs", handlers.ListJobs)

	app.Post("/job", handlers.CreateJob)

	app.Delete("/jobs/:id", handlers.DeleteJob)

	app.Put("/jobs/:id", handlers.UpdateJob)
}
