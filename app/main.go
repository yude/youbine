package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	_ "embed"

	"github.com/golang-jwt/jwt/v4"

	controllers "github.com/yude/youbine/controllers"
	"github.com/yude/youbine/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/template/html"
)

//go:embed public/*
var publicfs embed.FS

func main() {
	if os.Getenv("ADMIN_PASSWORD") == "" {
		log.Fatal("環境変数 ADMIN_PASSWORD を設定してから起動してください。")
	}

	database.Init()

	engine := html.NewFileSystem(http.FS(publicfs), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("public/login", fiber.Map{
			"notice": "ログインしてください。",
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("public/index", fiber.Map{
			"notice": "メッセージを送っていただけますと幸いです♪",
		})
	})
	app.Post("/", timeout.New(post_message, 60*time.Second))

	app.Post("/admin/login", timeout.New(controllers.Login, 60*time.Second))
	app.Get("/admin", admin)

	log.Fatal(app.Listen(":3000"))
}

func admin(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)

	if err != nil || !token.Valid {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "Unauthorized.",
		})
	}

	messages := database.ReturnMessage()

	return ctx.Render("public/admin", fiber.Map{
		"messages": messages,
	})
}
