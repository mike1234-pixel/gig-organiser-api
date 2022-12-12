package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mike1234-pixel/gig-organiser-api/database"
	"github.com/mike1234-pixel/gig-organiser-api/models"
)

func ListActions(c *fiber.Ctx) error {
	actions := []models.Action{}

	database.DB.Db.Find(&actions)

	return c.Status(200).JSON(actions)
}

func GetActions(c *fiber.Ctx) error {
	// Parse the userID parameter from the URL query string
	userID := c.Query("userID")

	// Query the actions table for actions with the given userID
	actions := []models.Action{}
	database.DB.Db.Where("user_id = ?", userID).Find(&actions)

	// Return the list of actions as a JSON array
	return c.JSON(actions)
}

func CreateAction(c *fiber.Ctx) error {
	action := new(models.Action)
	if err := c.BodyParser(action); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&action)

	return c.Status(200).JSON(action)
}

func UpdateAction(c *fiber.Ctx) error {
	id := c.Params("id")

	action := new(models.Action)
	if err := database.DB.Db.Where("id = ?", id).First(&action).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Action not found",
		})
	}

	if err := c.BodyParser(action); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Db.Save(&action).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(action)
}

func DeleteAction(c *fiber.Ctx) error {
	id := c.Params("id")

	action := new(models.Action)
	if err := database.DB.Db.Where("id = ?", id).First(&action).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Action not found",
		})
	}

	if err := database.DB.Db.Unscoped().Delete(&action).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
