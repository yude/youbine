package main

import (
	"log"
	"time"

	"github.com/yude/youbine/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func main() {
	database.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("public/index.html", fiber.Map{})
	})
	app.Post("/post", timeout.New(post_message, 60*time.Second))

	log.Fatal(app.Listen(":3000"))
}
