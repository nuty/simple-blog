package tests

import (
	"fmt"
	"bytes"
	"encoding/json"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/nuty/simple-blog/models"
	"github.com/nuty/simple-blog/utils"
)


func TestCreateArticle(t *testing.T) {
	app := SetupApp()
	randomStr, _ := utils.GenerateRandomString(12)

	article := models.Article{
		Slug:        randomStr,
		Title:       "title",
		Description: "description",
		Content:     "content",
	}
	articleData, _ := json.Marshal(article)

	req := httptest.NewRequest("POST", "/articles", bytes.NewReader(articleData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)


	fmt.Println(response)
	articleObject := response["article"].(map[string]interface{})

	assert.Equal(t, randomStr, articleObject["slug"])
	assert.Equal(t, "title", articleObject["title"])
	assert.Equal(t, "description", articleObject["description"])
	assert.Equal(t, "content", articleObject["content"])
}



func TestGetArticles(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("GET", "/articles", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)

	assert.Contains(t, response, "articles")
}
