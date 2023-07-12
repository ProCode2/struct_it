package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/procode2/structio/database"
	"github.com/procode2/structio/helpers"
	"github.com/procode2/structio/models"
)

func HandleGetSingin(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect("/")
	}

	return c.Render("auth/signin", nil)
}

func HandlePostSingin(c *fiber.Ctx) error {
	var user models.User

	email := c.FormValue("email")
	password := c.FormValue("password")
	name := c.FormValue("name")

	if email == "" || password == "" || name == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not enough data provides",
		})
	}

	user.Email = email
	user.Password = password
	user.Name = name

	newUser, err := database.Db.CreateNewUser(&user)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	jwtString, err := helpers.GetJWTKey(newUser)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "user",
		Value:   jwtString,
		Expires: time.Now().Add(24 * time.Hour),
	})

	return c.Redirect("/")
}

func HandleGetLogin(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.RedirectBack("/")
	}
	return c.Render("auth/login", nil)
}

func HandlePostLogin(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not enough data provides",
		})
	}

	user, err := database.Db.GetUserByEmail(email)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	fmt.Println(user)

	jwtString, err := helpers.GetJWTKey(user)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "user",
		Value:   jwtString,
		Expires: time.Now().Add(24 * time.Hour),
	})

	return c.Redirect("/")
}

func HandleDeleteLogout(c *fiber.Ctx) error {
	return c.SendString("Hello this is auth user")
}
