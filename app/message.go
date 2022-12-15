package main

import (
	_ "embed"

	"github.com/gofiber/fiber/v2"
	"github.com/yude/youbine/database"
)

//go:embed public/post.html
var post_html string

func post_message(c *fiber.Ctx) error {
	value := c.FormValue("value")

	if value == "" {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(post_html)
	} else {
		database.AddMessage(value, c.IP()+":"+c.Port())

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(post_html)
	}
}
