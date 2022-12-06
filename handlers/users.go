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

func GetUser(c *fiber.Ctx) error {
	// Get the user's email and password from the request
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Find the user in the database by their email
	user := models.User{}
	if err := database.DB.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Compare the user's password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// Return the user if the password is correct
	return c.Status(200).JSON(user)
}
