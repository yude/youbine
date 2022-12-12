package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yude/youbine/database"
)

func post_message(c *fiber.Ctx) error {
	value := c.FormValue("value")

	database.AddMessage(value, c.IP()+":"+c.Port())

	return c.Render("public/index.html", fiber.Map{
		"Message": `<div class="alert alert-success" role="alert">
						<p>✅ 送信しました。メッセージありがとうございます♪</p>
					</div>`,
	})
}
