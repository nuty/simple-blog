package handler

import (
	"encoding/json"
	"fmt"
	"github.com/nuty/simple-blog/models"
	"github.com/nuty/simple-blog/database"
	"github.com/nuty/simple-blog/providers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
	"context"
	"log"
	// "strconv"
)

var Rdb *redis.Client

type CommentHandler struct{}

func (handler CommentHandler) List(ctx *fiber.Ctx) error {
	articleID := ctx.Params("article_id")
	sortOrder := ctx.Query("sort", "desc")

	if sortOrder != "asc" && sortOrder != "desc" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid sort parameter, must be 'asc' or 'desc'",
		})
	}

	cacheKey := fmt.Sprintf("comments:article:%s:sort:%s", articleID, sortOrder)
	var comments []models.Comment

	val, err := providers.Rdb.Get(context.Background(), cacheKey).Result()
	if err == redis.Nil {
		// Do nothing wait for database fetch
	} else if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Redis error: " + err.Error(),
		})
	} else {
		// Cache hit, unmarshal and return cached comments
		if err := json.Unmarshal([]byte(val), &comments); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse cached comments",
			})
		}
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"comments": comments,
		})
	}

	query := database.DB.Where("article_id = ? AND parent_comment_id IS NULL", articleID).
		Preload("Children").
		Preload("Children.Children").
		Preload("Children.Children.Children") // 加载三级评论

	if sortOrder == "asc" {
		query = query.Order("created_at ASC")
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Find(&comments).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get comments",
		})
	}

	cacheData, err := json.Marshal(comments)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to serialize comments",
		})
	}

	if err := providers.Rdb.Set(context.Background(), cacheKey, cacheData, time.Hour).Err(); err != nil {
		log.Printf("Failed to cache comments for article %s: %v", articleID, err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"comments": comments,
	})
}

func (handler CommentHandler) Create(ctx *fiber.Ctx) error {
	var comment models.Comment
	if err := ctx.BodyParser(&comment); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	comment.CreatedAt = time.Now()
	if err := database.DB.Create(&comment).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add comment",
		})
	}

	cacheKey := fmt.Sprintf("comments:article:%d:sort:desc", comment.ArticleID)
	providers.Rdb.Del(context.Background(), cacheKey)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Comment posted successfully",
		"comment": comment,
	})
}

func (handler CommentHandler) Delete(ctx *fiber.Ctx) error {
	commentID := ctx.Params("comment_id")
	var comment models.Comment
	if err := database.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Comment not found",
		})
	}
	if err := database.DB.Delete(&comment).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete comment",
		})
	}

	// Clear the cache for the specific article, as we just deleted a comment
	cacheKey := fmt.Sprintf("comments:article:%d:sort:desc", comment.ArticleID) // 使用 fmt.Sprintf
	providers.Rdb.Del(context.Background(), cacheKey)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "deleted successfully",
	})
}
