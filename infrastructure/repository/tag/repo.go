package tag

import "github.com/zakiyalmaya/assetfindr-assignment/model"

type TagRepository interface {
	Create(tag *model.Tag) error
	GetByLabel(label string) (*model.Tag, error)
}