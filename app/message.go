package main

import (
	_ "embed"

	"github.com/gofiber/fiber/v2"
	"github.com/yude/youbine/database"
)

func post_message(c *fiber.Ctx) error {
	value := c.FormValue("value")

	if value != "" {
		database.AddMessage(value, c.IP()+":"+c.Port())
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Render("public/index", fiber.Map{
		"notice": "メッセージありがとうございました♪",
	})
}
