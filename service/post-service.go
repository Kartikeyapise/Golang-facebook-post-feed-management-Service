package service

import (
	"errors"
	"github.com/kartikeya/sample_app/entity"
	"github.com/kartikeya/sample_app/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(postRepo repository.PostRepository) PostService {
	repo = postRepo
	return &service{}
}

func (s service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("The Post is Empty")
	}
	if post.Title == "" {
		return errors.New("The Post Title is Empty")
	}
	return nil
}

func (service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int()
	return repo.Save(post)
}

func (service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
