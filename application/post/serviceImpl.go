package post

import (
	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
)

type postSvcImpl struct {
	repo *repository.Repository
}

func NewPostService(repo *repository.Repository) PostService {
	return &postSvcImpl{repo: repo}
}

func (p *postSvcImpl) Create(request *model.CreatePostRequest) error {
	return p.repo.Post.Create(post)
}

func (p *postSvcImpl) GetAll() ([]*model.Post, error) {
	return p.repo.Post.GetAll()
}

func (p *postSvcImpl) GetByID(id int) (*model.Post, error) {
	return p.repo.Post.GetByID(id)
}

func (p *postSvcImpl) Update(post *model.Post) error {
	return p.repo.Post.Update(post)
}

func (p *postSvcImpl) Delete(id int) error {
	return p.repo.Post.Delete(id)
}