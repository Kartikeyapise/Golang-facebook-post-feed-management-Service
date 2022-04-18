package repository

import (
	"github.com/kartikeya/sample_app/entity"
	"gorm.io/gorm"
	"log"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	err := r.DB.Create(post).Error
	log.Println("yooo", post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	err := r.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func NewPostgresRepository(db *gorm.DB) PostRepository {
	return &repo{db}
}
