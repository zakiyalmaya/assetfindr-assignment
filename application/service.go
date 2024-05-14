package application

import "github.com/zakiyalmaya/assetfindr-assignment/model"

type Service interface {
	Create(request *model.CreatePostRequest) error
	GetAll() ([]*model.Post, error)
	GetByID(id int) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id int) error
}