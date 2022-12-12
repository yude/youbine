package main

import (
	"log"
	"net/http"
	"time"

	_ "embed"

	"github.com/markbates/pkger"
	"github.com/yude/youbine/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/template/html"
)

func main() {
	database.Init()

	engine := html.NewFileSystem(pkger.Dir("/public"), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Provide a minimal config
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./public"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index.html", fiber.Map{})
	})
	app.Post("/post", timeout.New(post_message, 60*time.Second))

	log.Fatal(app.Listen(":3000"))
}
