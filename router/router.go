package router

import (
	"github.com/nuty/simple-blog/handlers"
	"github.com/nuty/simple-blog/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Initalize(router *fiber.App) {
	router.Use(middlewares.Json)
	api := router.Group("api/v1/")

	articleHandler := new(handler.ArticleHandler)
	commentHandler := new(handler.CommentHandler)


	api.Get("/articles", articleHandler.List)
	api.Post("/articles", articleHandler.Create)
	
	api.Get("/articles/:article_id/comments", commentHandler.List)

	api.Post("/comments", commentHandler.Create)
	api.Delete("/comments/:comment_id", commentHandler.Delete)
}