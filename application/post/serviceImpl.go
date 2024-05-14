package post

import (
	"log"

	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/gorm"
)

type postSvcImpl struct {
	repo *repository.Repository
}

func NewPostService(repo *repository.Repository) PostService {
	return &postSvcImpl{repo: repo}
}

func (p *postSvcImpl) Create(request *model.CreatePostRequest) error {
	tags := make([]*model.Tag, len(request.Tags))
	for i, label := range request.Tags {
		var tag *model.Tag
		tag, err := p.repo.Tag.GetByLabel(label)
		if err == gorm.ErrRecordNotFound {
			tag = &model.Tag{Label: label}
			if err := p.repo.Tag.Create(tag); err != nil {
				log.Println("error creating tag: ", err.Error())
				return err
			}
		} else if err != nil {
			log.Println("error getting tag: ", err.Error())
			return err
		}

		tags[i] = tag
	}

	return p.repo.Post.Create(&model.Post{
		Title:   request.Title,
		Content: request.Content,
		Tags:    tags,
	})
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