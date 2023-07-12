package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/procode2/structio/database"
	"github.com/procode2/structio/middlewares"
	"github.com/procode2/structio/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("Loaded dotenv %s\n", os.Getenv("DB_USER"))
	app := InitApp()
	database.Db.Init()
	app.Use(middlewares.InjectAuthUser)
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

func InitApp() *fiber.App {
	engine := InitViewEngine()
	app := fiber.New(fiber.Config{
		Views:             engine,
		ViewsLayout:       "layouts/main",
		PassLocalsToViews: true,
	})

	app.Static("/", "./static", fiber.Static{
		CacheDuration: 0,
	})

	return app
}

func InitViewEngine() *html.Engine {
	engine := html.New("./views", ".tmpl")
	engine.Reload(true)
	engine.AddFunc("css", func(name string) (res template.HTML) {
		filepath.Walk("static/css", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Name() == name {
				res = template.HTML(fmt.Sprintf("<link rel=\"stylesheet\" href=\"/css/%s\">", name))
			}

			return nil
		})

		return
	})

	return engine
}
