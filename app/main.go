package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	_ "embed"

	"github.com/yude/youbine/controllers"
	"github.com/yude/youbine/database"
	"github.com/yude/youbine/routes"
	"github.com/yude/youbine/webhook"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/template/html"
)

//go:embed static/*
var staticfs embed.FS

func main() {
	if os.Getenv("ADMIN_PASSWORD") == "" {
		log.Fatal("環境変数 ADMIN_PASSWORD を設定してから起動してください。")
	}
	if os.Getenv("TZ") == "" {
		os.Setenv("TZ", "Asia/Tokyo")
	}

	database.Init()
	webhook.Initialize()

	engine := html.NewFileSystem(http.FS(staticfs), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("static/index", fiber.Map{
			"notice": "メッセージを送っていただけますと幸いです♪",
		})
	})
	app.Post("/", timeout.New(controllers.Post, 60*time.Second))

	app.Get("/admin", routes.Admin)
	app.Post("/login", timeout.New(controllers.Login, 60*time.Second))

	log.Fatal(app.Listen(":3000"))
}
