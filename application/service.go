package application

import "github.com/zakiyalmaya/assetfindr-assignment/model"

type Service interface {
	Create(request *model.PostRequest) error
	GetAll() ([]*model.Post, error)
	GetByID(id int) (*model.Post, error)
	Update(id int, request *model.PostRequest) error
	Delete(id int) error
}