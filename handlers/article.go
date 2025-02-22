package handler

import (
	"github.com/nuty/simple-blog/models"
	"github.com/nuty/simple-blog/database"

	"net/http"
	"github.com/gofiber/fiber/v2"
)


type ArticleHandler struct{}

func (handler ArticleHandler) Create(ctx *fiber.Ctx) error {
	var article models.Article
	if err := ctx.BodyParser(&article); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if err := database.DB.Create(&article).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create article",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Article created successfully",
		"article": article,
	})
}


func (handler ArticleHandler) List(ctx *fiber.Ctx) error {
	sortOrder := ctx.Query("sort", "desc")
	var articles []models.Article
	if err := database.DB.Order("created_at " + sortOrder).
		Find(&articles).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get articles",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"articles":    articles,
	})
}


