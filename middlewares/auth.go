package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/procode2/structio/database"
	"github.com/procode2/structio/helpers"
)

func InjectAuthUser(c *fiber.Ctx) error {
	fmt.Printf("I am being called by %+v\nlocal: %s\n", c.OriginalURL(), c.Locals("user"))

	if c.Locals("user") != nil {
		return c.Next()
	}

	jwtString := c.Cookies("user", "not_found")

	if jwtString == "not_found" {
		return c.Next()
	}

	claim, err := helpers.ValiadateTokenString(jwtString)

	if err != nil {
		return c.Next()
	}

	user, err := database.Db.GetUserById(claim.Id)
	if err != nil {
		return c.Next()
	}
	c.Locals("user", user)

	return c.Next()
}
