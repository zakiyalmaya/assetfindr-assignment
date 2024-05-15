package application

import (
	"log"
	"sync"

	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/gorm"
)

type serviceImpl struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) Service {
	return &serviceImpl{repo: repo}
}

func (p *serviceImpl) Create(request *model.CreatePostRequest) error {
	var wg sync.WaitGroup
	tags := make([]*model.Tag, len(request.Tags))
	errChan := make(chan error, len(request.Tags))

	for i, label := range request.Tags {
		wg.Add(1)
		go func(i int, label string) {
			defer wg.Done()

			tag, err := p.repo.Tag.GetByLabel(label)
			if err == gorm.ErrRecordNotFound {
				tag = &model.Tag{Label: label}
				if err := p.repo.Tag.Create(tag); err != nil {
					errChan <- err
					return
				}
			} else if err != nil {
				errChan <- err
				return
			}

			tags[i] = tag
		}(i, label)
	}
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			log.Println("error creating or getting tag: ", err.Error())
			return err
		}
	}

	return p.repo.Post.Create(&model.Post{
		Title:   request.Title,
		Content: request.Content,
		Tags:    tags,
	})
}

func (p *serviceImpl) GetAll() ([]*model.Post, error) {
	return p.repo.Post.GetAll()
}

func (p *serviceImpl) GetByID(id int) (*model.Post, error) {
	return p.repo.Post.GetByID(id)
}

func (p *serviceImpl) Update(post *model.Post) error {
	return p.repo.Post.Update(post)
}

func (p *serviceImpl) Delete(id int) error {
	return p.repo.Post.Delete(id)
}
