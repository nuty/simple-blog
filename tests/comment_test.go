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

// // TestCreateComment 测试创建评论的接口
// func TestCreateComment(t *testing.T) {
// 	app := SetupApp()

// 	// 创建评论数据
// 	comment := models.Comment{
// 		ArticleID: 1, // 固定文章ID为1
// 		Username:  "Test User",
// 		Email:     "testuser@example.com",
// 		Content:   "This is a test comment",
// 	}

// 	commentData, _ := json.Marshal(comment)

// 	// 模拟 POST 请求创建评论
// 	req := httptest.NewRequest("POST", "/comments", bytes.NewReader(commentData))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 断言状态码
// 	assert.Equal(t, 200, resp.StatusCode)

// 	// 解析响应 body
// 	var response map[string]interface{}
// 	_ = json.NewDecoder(resp.Body).Decode(&response)

// 	// 验证响应内容
// 	assert.Equal(t, "Comment posted successfully", response["message"])

// 	// 验证评论字段
// 	commentObject := response["comment"].(map[string]interface{})
// 	assert.Equal(t, "Test User", commentObject["username"])
// 	assert.Equal(t, "testuser@example.com", commentObject["email"])
// 	assert.Equal(t, "This is a test comment", commentObject["content"])
// }

// // TestListComments 测试获取评论列表接口
// func TestListComments(t *testing.T) {
// 	app := SetupApp()

// 	// 模拟请求获取评论列表
// 	req := httptest.NewRequest("GET", "/articles/1/comments", nil)
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 断言状态码
// 	assert.Equal(t, 200, resp.StatusCode)

// 	// 解析响应 body
// 	var response map[string]interface{}
// 	_ = json.NewDecoder(resp.Body).Decode(&response)

// 	// 验证响应结构
// 	assert.Contains(t, response, "comments")
// 	assert.NotNil(t, response["comments"])

// 	// 验证评论列表不为空
// 	comments := response["comments"].([]interface{})
// 	assert.NotEmpty(t, comments, "No comments found")
// }

// // TestDeleteComment 测试删除评论接口
// func TestDeleteComment(t *testing.T) {
// 	app := SetupApp()

// 	// 创建评论数据
// 	comment := models.Comment{
// 		ArticleID: 1, // 固定文章ID为1
// 		Username:  "Delete User",
// 		Email:     "deleteuser@example.com",
// 		Content:   "This comment will be deleted",
// 	}

// 	commentData, _ := json.Marshal(comment)

// 	// 模拟 POST 请求创建评论
// 	req := httptest.NewRequest("POST", "/comments", bytes.NewReader(commentData))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 解析响应 body
// 	var response map[string]interface{}
// 	_ = json.NewDecoder(resp.Body).Decode(&response)

// 	// 获取刚创建的评论ID
// 	commentID := uint(response["comment"].(map[string]interface{})["id"].(float64))

// 	// 模拟删除评论
// 	deleteReq := httptest.NewRequest("DELETE", fmt.Sprintf("/comments/%d", commentID), nil)
// 	deleteResp, err := app.Test(deleteReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 断言状态码
// 	assert.Equal(t, 200, deleteResp.StatusCode)

// 	// 验证删除后的响应内容
// 	var deleteResponse map[string]interface{}
// 	_ = json.NewDecoder(deleteResp.Body).Decode(&deleteResponse)
// 	assert.Equal(t, "deleted successfully", deleteResponse["message"])

// 	// 尝试从数据库获取已删除评论
// 	var deletedComment models.Comment


// 	err = database.DB.First(&deletedComment, commentID).Error
// 	assert.Error(t, err, "Expected error when fetching deleted comment from DB")
// }
// func TestDeleteCommentAndChildren(t *testing.T) {
// 	// Step 1: 创建一个父评论
// 	parentComment := models.Comment{
// 		ArticleID: 1,
// 		Username:  "user1",
// 		Email:     "user1@example.com",
// 		Content:   "This is a parent comment",
// 	}
// 	database.DB.Create(&parentComment)

// 	// Step 2: 创建两个子评论
// 	childComment1 := models.Comment{
// 		ArticleID:      1,
// 		ParentCommentID: &parentComment.ID,
// 		Username:       "user2",
// 		Email:          "user2@example.com",
// 		Content:        "This is the first child comment",
// 	}
// 	database.DB.Create(&childComment1)

// 	childComment2 := models.Comment{
// 		ArticleID:      1,
// 		ParentCommentID: &parentComment.ID,
// 		Username:       "user3",
// 		Email:          "user3@example.com",
// 		Content:        "This is the second child comment",
// 	}
// 	database.DB.Create(&childComment2)

// 	// Step 3: 发送删除请求
// 	app := SetupApp()
// 	req := httptest.NewRequest("DELETE", fmt.Sprintf("/comments/%d", parentComment.ID), nil)
// 	resp, _ := app.Test(req)

// 	// Step 4: 检查删除是否成功
// 	assert.Equal(t, 200, resp.StatusCode)

// 	// Step 5: 验证评论和子评论是否被删除
// 	var parent models.Comment
// 	err := database.DB.First(&parent, parentComment.ID).Error
// 	assert.Equal(t, gorm.ErrRecordNotFound, err)

// 	var child1 models.Comment
// 	err = database.DB.First(&child1, childComment1.ID).Error
// 	assert.Equal(t, gorm.ErrRecordNotFound, err)

// 	var child2 models.Comment
// 	err = database.DB.First(&child2, childComment2.ID).Error
// 	assert.Equal(t, gorm.ErrRecordNotFound, err)
// }


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
