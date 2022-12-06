package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mike1234-pixel/gig-organiser-api/database"
	"github.com/mike1234-pixel/gig-organiser-api/models"
	"golang.org/x/crypto/bcrypt"
)

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Db.Find(&users)

	return c.Status(200).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Check if a user with the submitted email already exists
	existingUser := models.User{}
	if err := database.DB.Db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "A user with the submitted email already exists",
		})
	}

	// Hash the user's password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Save the hashed password in the database
	user.Password = string(hashedPassword)
	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}
