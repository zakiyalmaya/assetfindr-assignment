package post

import "github.com/zakiyalmaya/assetfindr-assignment/model"

type PostRepository interface {
	Create(post *model.Post) error
	GetAll() ([]*model.Post, error)
	GetByID(id int) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id int) error
}
