package post

import (
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/gorm"
)

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=PostRepository.go
type PostRepository interface {
	Create(post *model.Post, tx ...*gorm.DB) error
	GetAll() ([]*model.Post, error)
	GetByID(id int, tx ...*gorm.DB) (*model.Post, error)
	Update(post *model.Post, tx ...*gorm.DB) error
	Delete(id int, tx ...*gorm.DB) error
	Assosiate(post *model.Post, tags []*model.Tag, tx ...*gorm.DB) error
}
