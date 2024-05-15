package application

import (
	"log"

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

func (p *serviceImpl) Create(request *model.PostRequest) error {
	tx := p.repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tags, err := p.getOrCreateTags(request.Tags, tx)
	if err != nil {
		log.Println("error fetching tags: ", err.Error())
		tx.Rollback()
		return err
	}

	if err := p.repo.Post.Create(&model.Post{
		Title:   request.Title,
		Content: request.Content,
		Tags:    tags,
	}, tx); err != nil {
		log.Println("error creating post: ", err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("error committing transaction: ", err.Error())
		tx.Rollback()
		return err
	}

	return nil
}

func (p *serviceImpl) GetAll() ([]*model.Post, error) {
	return p.repo.Post.GetAll()
}

func (p *serviceImpl) GetByID(id int) (*model.Post, error) {
	return p.repo.Post.GetByID(id)
}

func (p *serviceImpl) Update(id int, request *model.PostRequest) error {
	tx := p.repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	existingPost, err := p.repo.Post.GetByID(id, tx)
	if err != nil {
		log.Println("error getting post: ", err.Error())
		tx.Rollback()
		return err
	}

	tags, err := p.getOrCreateTags(request.Tags, tx)
	if err != nil {
		log.Println("error fetching tags: ", err.Error())
		tx.Rollback()
		return err
	}

	updatedPost := &model.Post{
		ID:      existingPost.ID,
		Title:   request.Title,
		Content: request.Content,
		Tags:    tags,
	}
	if err := p.repo.Post.Update(updatedPost, tx); err != nil {
		log.Println("error updating post: ", err.Error())
		tx.Rollback()
		return err
	}

	if err := p.repo.Post.Assosiate(updatedPost, tags, tx); err != nil {
		log.Println("error assosiating post: ", err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("error committing transaction: ", err.Error())
		tx.Rollback()
		return err
	}

	return nil
}

func (p *serviceImpl) Delete(id int) error {
	tx := p.repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err := p.repo.Post.GetByID(id, tx)
	if err != nil {
		log.Println("error getting post: ", err.Error())
		tx.Rollback()
		return err
	}
	
	if err := p.repo.Post.Delete(id, tx); err != nil {
		log.Println("error deleting post: ", err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("error committing transaction: ", err.Error())
		tx.Rollback()
		return err
	}
	
	return nil
}

func (p *serviceImpl) getOrCreateTags(requestTags []string, tx *gorm.DB) ([]*model.Tag, error) {
	tags := make([]*model.Tag, len(requestTags))
	for i, label := range requestTags {
		tag, err := p.repo.Tag.GetOrCreate(&model.Tag{Label: label}, tx)
		if err != nil {
			return nil, err
		}

		tags[i] = tag
	}

	return tags, nil
}
