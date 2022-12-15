package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	utils "github.com/yude/youbine/utils"
)

func Login(ctx *fiber.Ctx) error {
	value := ctx.FormValue("value")

	origin_password := []byte(utils.GetEnv("ADMIN_PASSWORD", "$2a$12$JBJza9iN0StqcAICG8xuv.ffEMyU2okVgOGos0CfVFRj6W/GyrMDi"))

	err := bcrypt.CompareHashAndPassword(origin_password, []byte(value))
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	payload := jwt.StandardClaims{
		Subject:   "admin",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.Redirect("/admin")
}
