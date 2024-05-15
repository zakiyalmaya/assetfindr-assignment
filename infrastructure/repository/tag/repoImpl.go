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

func (t *tagRepoImpl) GetOrCreate(tag *model.Tag, tx *gorm.DB) (*model.Tag, error) {
	if err := tx.FirstOrCreate(&tag, &model.Tag{Label: tag.Label}).Error; err != nil {
		log.Println("error creating or getting tag: ", err.Error())
		return nil, err
	}

	return tag, nil
}