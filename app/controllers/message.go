package controllers

import (
	_ "embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yude/youbine/database"
)

func Post(c *fiber.Ctx) error {
	value := c.FormValue("value")
	client_ip := c.IP()

	notice := "メッセージありがとうございました♪"

	if c.GetRespHeader("X-Forwarded-For") != "" {
		client_ip = c.GetRespHeader("X-Forwarded-For")
	}
	if c.GetRespHeader("CF-Connecting-IP") != "" {
		client_ip = c.GetRespHeader("CF-Connecting-IP")
	}

	if value != "" {
		database.AddMessage(value, client_ip+":"+c.Port())
	} else {
		notice += " 送れているかは別として・・・"
	}

	log.Print("New message: " + value + " from " + client_ip + ":" + c.Port())

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Render("static/index", fiber.Map{
		"notice": notice,
	})
}
