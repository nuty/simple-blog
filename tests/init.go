package tests

import (
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nuty/simple-blog/handlers"
	"github.com/nuty/simple-blog/database"
	"github.com/nuty/simple-blog/providers"
	"github.com/nuty/simple-blog/config"
)

func SetupApp() *fiber.App {
	app := fiber.New()	
	dir, _ := os.Getwd()

	fmt.Println(dir)


	config, _ := config.LoadConfig("../config/config.toml")
	database.ConnectDB(config)
	providers.InitRedis()

	articleHandler := handler.ArticleHandler{}
	commentHandler := handler.CommentHandler{}

	app.Get("/articles", articleHandler.List)
	app.Post("/articles", articleHandler.Create)
	
	app.Get("/articles/:article_id/comments", commentHandler.List)

	app.Post("/comments", commentHandler.Create)
	app.Delete("/comments/:comment_id", commentHandler.Delete)
	return app
}