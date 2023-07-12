package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/procode2/structio/database"
	"github.com/procode2/structio/models"
)

func HandleGetHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func HandleGetExplorePaths(c *fiber.Ctx) error {
	search := c.Query("q")
	paths, err := database.Db.GetAllPaths(search)
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}

	return c.Render("paths/explore", fiber.Map{
		"Paths": paths,
	})
}

func HandlePostPath(c *fiber.Ctx) error {
	// path := &models.Path{}
	// if err := c.BodyParser(path); err != nil {
	// 	if err != nil {
	// 		c.Redirect("/")
	// 	}
	// }

	path := &models.Path{
		Title:       "This is a demo title",
		Description: "extraordinary Description",
		Tags:        []string{"tag1", "tag2"},
		Levels: []*models.Level{
			{
				LevelNo: 0,
				Bits: []*models.Bit{
					{
						Link:        "demo link",
						Description: "demo bit desc",
					},
				},
			},
		},
	}

	path, err := database.Db.CreateNewPath(path)
	if err != nil {
		c.Redirect("/")
	}
	fmt.Println("path")
	fmt.Println(path)
	return c.Render("paths/path", fiber.Map{
		"Path": path,
	})
}

func HandleGetPath(c *fiber.Ctx) error {
	pathId := c.Params("pathId")
	path, err := database.Db.GetPathById(pathId)
	if err != nil {
		c.Redirect("/")
	}

	return c.Render("paths/path", fiber.Map{
		"Path": path,
		"Js":   []string{"/build/component2.js"},
		"Css":  []string{"/build/assets/component2.css"},
	})
}

func HandleUpdatePath(c *fiber.Ctx) error {
	_ = c.Params("pathId")

	path := &models.Path{}
	if err := c.BodyParser(path); err != nil {
		if err != nil {
			c.Redirect("/")
		}
	}
	err := database.Db.UpdatePath(path)
	if err != nil {
		c.Redirect("/")
	}

	return c.Render("paths/path", fiber.Map{
		"Path": path,
	})
}
