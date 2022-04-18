package repository

import "github.com/kartikeya/sample_app/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
