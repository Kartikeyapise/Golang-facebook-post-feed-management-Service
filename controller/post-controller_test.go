package controller

import (
	"bytes"
	"encoding/json"
	"github.com/kartikeya/sample_app/entity"
	"github.com/kartikeya/sample_app/repository"
	"github.com/kartikeya/sample_app/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	DB                 *gorm.DB                  = repository.ConnectDatabase()
	PostRepositoryTest repository.PostRepository = repository.NewPostgresRepository(DB)
	postServiceTest    service.PostService       = service.NewPostService(PostRepositoryTest)
	postControllerTest PostController            = NewPostController(postServiceTest)
)

func cleanDatabase() {
	DB.Delete(entity.Post{}, "title LIKE ?", "%%")
}

func addSamplePost() {
	post := entity.Post{
		Id:    1,
		Title: "title",
		Text:  "text",
	}
	DB.Create(&post)
}
func TestGetPosts(t *testing.T) {
	cleanDatabase()
	addSamplePost()
	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/posts", nil)

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(postControllerTest.GetPosts)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	//Decode the HTTP response
	var posts []entity.Post
	json.NewDecoder(response.Body).Decode(&posts)

	//Assert HTTP response
	assert.NotNil(t, posts)
	assert.Equal(t, "title", posts[0].Title)
	assert.Equal(t, "text", posts[0].Text)
}

func TestAddPost(t *testing.T) {
	//Create a new HTTP POST request
	var jsonData = []byte(`{"title":"Title 1", "text":"Text 1"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonData))

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(postControllerTest.AddPost)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	//Decode the HTTP response
	var post entity.Post
	json.NewDecoder(response.Body).Decode(&post)

	//Assert HTTP response
	assert.NotNil(t, post)
	assert.Equal(t, "Title 1", post.Title)
	assert.Equal(t, "Text 1", post.Text)
}
