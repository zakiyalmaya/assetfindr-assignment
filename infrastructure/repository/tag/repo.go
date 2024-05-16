package tag

import (
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/gorm"
)

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=TagRepository.go
type TagRepository interface {
	GetOrCreate(tag *model.Tag, tx *gorm.DB) (*model.Tag, error)
}