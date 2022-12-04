package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mike1234-pixel/gig-organiser-api/database"
	"github.com/mike1234-pixel/gig-organiser-api/models"
)

func ListJobs(c *fiber.Ctx) error {
	jobs := []models.Job{}

	database.DB.Db.Find(&jobs)

	return c.Status(200).JSON(jobs)
}

func CreateJob(c *fiber.Ctx) error {
	job := new(models.Job)
	if err := c.BodyParser(job); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&job)

	return c.Status(200).JSON(job)
}

func DeleteJob(c *fiber.Ctx) error {
	id := c.Params("id")

	job := new(models.Job)
	if err := database.DB.Db.Where("id = ?", id).First(&job).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Job not found",
		})
	}

	if err := database.DB.Db.Unscoped().Delete(&job).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateJob(c *fiber.Ctx) error {
	id := c.Params("id")

	job := new(models.Job)
	if err := database.DB.Db.Where("id = ?", id).First(&job).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Job not found",
		})
	}

	if err := c.BodyParser(job); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Db.Save(&job).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(job)
}
