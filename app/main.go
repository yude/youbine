package main

import (
	"log"
	"time"

	_ "embed"

	"github.com/yude/youbine/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

//go:embed public/index.html
var index_html string

func main() {
	database.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(index_html)
	})
	app.Post("/post", timeout.New(post_message, 60*time.Second))

	log.Fatal(app.Listen(":3000"))
}
