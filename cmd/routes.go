package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mike1234-pixel/gig-organiser-api/handlers"
)

func setupRoutes(app *fiber.App) {
	// Use the CORS middleware to enable CORS for fe
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://main.d283u7qcw25b1m.amplifyapp.com",
		AllowCredentials: true,
	}))

	// users

	app.Get("/users", handlers.ListUsers)

	app.Post("/create-user", handlers.CreateUser)

	app.Post("/user", handlers.GetUser)

	app.Delete("/user/:id", handlers.DeleteUser)

	// jobs

	app.Get("/all-jobs", handlers.ListJobs)

	app.Get("/jobs", handlers.GetJobs)

	app.Post("/job", handlers.CreateJob)

	app.Delete("/jobs/:id", handlers.DeleteJob)

	app.Put("/jobs/:id", handlers.UpdateJob)

	// actions

	app.Get("/all-actions", handlers.ListActions)

	app.Get("/actions", handlers.GetActions)

	app.Post("/action", handlers.CreateAction)

	app.Put("/actions/:id", handlers.UpdateAction)

	app.Delete("/actions/:id", handlers.DeleteAction)
}
