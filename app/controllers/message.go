package controllers

import (
	_ "embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yude/youbine/database"
)

func Post(c *fiber.Ctx) error {
	value := c.FormValue("value")

	notice := "メッセージありがとうございました♪"

	if value != "" {
		database.AddMessage(value, c.IP()+":"+c.Port())
	} else {
		notice += " 送れているかは別として・・・"
	}

	log.Print("New message: " + value + " from " + c.IP() + ":" + c.Port())

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Render("static/index", fiber.Map{
		"notice": notice,
	})
}
