package tag

import (
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/gorm"
)

type TagRepository interface {
	GetOrCreate(tag *model.Tag, tx *gorm.DB) (*model.Tag, error)
}