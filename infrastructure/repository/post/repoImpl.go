package post

import (
	"log"

	"gorm.io/gorm"

	"github.com/zakiyalmaya/assetfindr-assignment/model"
)

type postRepoImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepoImpl{db: db}
}

func (p *postRepoImpl) Create(post *model.Post, tx ...*gorm.DB) error {
	db := p.db
	if len(tx) > 0 {
		db = tx[0]
	}
	
	if err := db.Create(post).Error; err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (p *postRepoImpl) GetAll() ([]*model.Post, error) {
	var posts []*model.Post
	if err := p.db.Preload("Tags").Find(&posts).Error; err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return posts, nil
}

func (p *postRepoImpl) GetByID(id int, tx ...*gorm.DB) (*model.Post, error) {
	db := p.db
	if len(tx) > 0 {
		db = tx[0]
	}

	var post *model.Post
	if err := db.Preload("Tags").First(&post, id).Error; err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return post, nil
}

func (p *postRepoImpl) Update(post *model.Post, tx ...*gorm.DB) error {
	db := p.db
	if len(tx) > 0 {
		db = tx[0]
	}

	if err := db.Save(post).Error; err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (p *postRepoImpl) Delete(id int) error {
	if err := p.db.Delete(&model.Post{}, id).Error; err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (p *postRepoImpl) Assosiate(post *model.Post, tags []*model.Tag, tx ...*gorm.DB) error {
	db := p.db
	if len(tx) > 0 {
		db = tx[0]
	}

	if err := db.Model(post).Association("Tags").Replace(tags); err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}