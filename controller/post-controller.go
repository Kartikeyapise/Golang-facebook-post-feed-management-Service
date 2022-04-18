package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kartikeya/sample_app/entity"
	"github.com/kartikeya/sample_app/errors"
	"github.com/kartikeya/sample_app/service"
	"log"
	"net/http"
)

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (c controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting posts"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
	return
}

func (c controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	resp.Header().Set("Content-type", "application/json")
	fmt.Println("reqBOdy::::", req.Body)
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	log.Println(post)
	res_post, err2 := postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Cannot Save the post in database"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(res_post)
	return
}

//
//func (controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
//	resp.Header().Set("Content-type", "application/json")
//	posts, err := postService.FindAll()
//	if err != nil {
//		resp.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting posts"})
//		return
//	}
//	resp.WriteHeader(http.StatusOK)
//	json.NewEncoder(resp).Encode(posts)
//	return
//}
//
//func (controller) AddPost(PostController) (resp http.ResponseWriter, req *http.Request) {
//	var post entity.Post
//	resp.Header().Set("Content-type", "application/json")
//	fmt.Println("reqBOdy::::", req.Body)
//	err := json.NewDecoder(req.Body).Decode(&post)
//	if err != nil {
//		resp.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
//		return
//	}
//	err1 := postService.Validate(&post)
//	if err1 != nil {
//		resp.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
//		return
//	}
//	log.Println(post)
//	res_post, err2 := postService.Create(&post)
//	if err2 != nil {
//		resp.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Cannot Save the post in database"})
//		return
//	}
//	resp.WriteHeader(http.StatusOK)
//	json.NewEncoder(resp).Encode(res_post)
//	return
//}
