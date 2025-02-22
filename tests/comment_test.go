package tests

import (
	"bytes"
	"encoding/json"
	"testing"
	"strconv"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	app := SetupApp()

	// 准备测试数据
	commentData := map[string]interface{}{
		"article_id":       1,
		"content":          "This is a test comment",
		"parent_comment_id": nil,
	}

	// 将数据转化为 JSON
	jsonData, err := json.Marshal(commentData)
	assert.NoError(t, err)

	// 发起 POST 请求
	req := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// 检查返回状态码
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 检查返回的数据
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Comment posted successfully", response["message"])
}

// 测试获取评论列表
func TestGetComments(t *testing.T) {
	app := SetupApp()

	// 创建一个文章的评论
	commentData := map[string]interface{}{
		"article_id":       1,
		"content":          "This is a test comment",
		"parent_comment_id": nil,
	}
	jsonData, err := json.Marshal(commentData)
	assert.NoError(t, err)

	// 发起 POST 请求，创建评论
	req := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	_, err = app.Test(req)
	assert.NoError(t, err)

	// 获取评论列表
	req = httptest.NewRequest(http.MethodGet, "/articles/1/comments", nil)
	resp, err := app.Test(req)

	// 检查返回状态码
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 检查返回的数据
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	comments, ok := response["comments"].([]interface{})
	assert.True(t, ok)
	assert.Greater(t, len(comments), 0) // 确保评论存在
}

// 测试删除评论
func TestDeleteComment(t *testing.T) {
	app := SetupApp()

	// 创建一个评论
	commentData := map[string]interface{}{
		"article_id":       1,
		"content":          "This is a test comment to delete",
		"parent_comment_id": nil,
	}
	jsonData, err := json.Marshal(commentData)
	assert.NoError(t, err)

	// 发起 POST 请求，创建评论
	req := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// 获取返回的评论 ID
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	commentID := response["comment"].(map[string]interface{})["id"].(float64)

	// 发起 DELETE 请求，删除评论
	req = httptest.NewRequest(http.MethodDelete, "/comments/"+strconv.Itoa(int(commentID)), nil)
	resp, err = app.Test(req)

	// 检查返回状态码
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 检查返回的数据
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "deleted successfully", response["message"])
}
