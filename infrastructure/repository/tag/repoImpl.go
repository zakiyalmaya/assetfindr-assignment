package tag

import (
	"log"

	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/gorm"
)
type tagRepoImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepoImpl{db: db}
}

func (t *tagRepoImpl) Create(tag *model.Tag) error {
	if err := t.db.Create(tag).Error; err != nil {
		log.Println("error creating tag: ", err.Error())
		return err
	}

	return nil
}

func (t *tagRepoImpl) GetByLabel(label string) (*model.Tag, error) {
	var tag model.Tag
	if err := t.db.Where("label = ?", label).First(&tag).Error; err != nil {
		log.Println("error getting tag by label: ", err.Error())
		return nil, err
	}

	return &tag, nil
}
