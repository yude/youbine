package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/yude/youbine/database"
)

func Admin(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)

	if err != nil || !token.Valid {
		return c.Render("static/login", fiber.Map{
			"notice": "このページを閲覧するには、ログインが必要です。",
		})
	}

	messages := database.ReturnMessage()

	return c.Render("static/admin", fiber.Map{
		"messages": messages,
	})
}
