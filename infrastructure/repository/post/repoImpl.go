package post

import (
	"gorm.io/gorm"

	"github.com/zakiyalmaya/assetfindr-assignment/model"
)

type postRepoImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepoImpl{db: db}
}

func (p *postRepoImpl) Create(post *model.Post) error {
	if err := p.db.Create(post).Error; err != nil {
		return err
	}

	return nil
}

func (p *postRepoImpl) GetAll() ([]*model.Post, error) {
	var posts []*model.Post
	if err := p.db.Preload("Tags").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *postRepoImpl) GetByID(id int) (*model.Post, error) {
	var post *model.Post
	if err := p.db.Preload("Tags").First(post, id).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepoImpl) Update(post *model.Post) error {
	if err := p.db.Save(post).Error; err != nil {
		return err
	}

	return nil
}

func (p *postRepoImpl) Delete(id int) error {
	if err := p.db.Delete(&model.Post{}, id).Error; err != nil {
		return err
	}

	return nil
}