package repository

import (
	"fmt"
	"log"

	"github.com/zakiyalmaya/assetfindr-assignment/config"
	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository/post"
	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository/tag"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db   *gorm.DB
	Post post.PostRepository
	Tag  tag.TagRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:   db,
		Post: post.NewPostRepository(db),
		Tag:  tag.NewTagRepository(db),
	}
}

func Database() *gorm.DB {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.PG_DB_HOST,
		config.PG_DB_PORT,
		config.PG_DB_USERNAME,
		config.PG_DB_PASSWORD,
		config.PG_DB_NAME,
		config.PG_DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Panicln("error connecting to database: ", err.Error())
		return nil
	}

	db.AutoMigrate(&model.Post{}, &model.Tag{})
	return db
}
