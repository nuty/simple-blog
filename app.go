package main

import (
	"log"
	"github.com/nuty/simple-blog/models"
	"github.com/nuty/simple-blog/database"
	"github.com/nuty/simple-blog/providers"
	"github.com/nuty/simple-blog/router"
	"github.com/nuty/simple-blog/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := config.LoadConfig("config/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	database.ConnectDB(config)
	database.DB.AutoMigrate(
		&models.Article{},
		&models.Comment{},
	)
	app := fiber.New()
	providers.InitRedis()
	router.Initalize(app)
	log.Fatal(app.Listen(":3001"))
}
