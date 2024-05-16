package application

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository"
	mocks "github.com/zakiyalmaya/assetfindr-assignment/mocks/infrastructure/repository"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mockPostRepo *mocks.MockPostRepository
	mockTagRepo  *mocks.MockTagRepository
	service      Service
	gormdb       *gorm.DB
)

func Setup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	sqldb, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}

	gormdb, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatalf("An error occurred while creating gorm db mock: %s", err)
	}

	mockPostRepo = mocks.NewMockPostRepository(mockCtl)
	mockTagRepo = mocks.NewMockTagRepository(mockCtl)
	service = NewService(&repository.Repository{
		DB:   gormdb,
		Post: mockPostRepo,
		Tag:  mockTagRepo,
	})
}

func TestGetAll(t *testing.T) {
	Setup(t)

	testCases := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when get all post then return success",
			mock: func() {
				mockPostRepo.EXPECT().GetAll().Return([]*model.Post{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Given error when get all post then return error",
			mock: func() {
				mockPostRepo.EXPECT().GetAll().Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			_, err := service.GetAll()
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	Setup(t)

	testCases := []struct {
		name    string
		id      int
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when get post by id then return success",
			id:   1,
			mock: func() {
				mockPostRepo.EXPECT().GetByID(1).Return(&model.Post{ID: 1}, nil)
			},
			wantErr: false,
		},
		{
			name: "Given error when get post by id then return error",
			id:   1,
			mock: func() {
				mockPostRepo.EXPECT().GetByID(1).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			_, err := service.GetByID(tc.id)
			if err != nil && !tc.wantErr {
				t.Fatal(err)
			}
		})
	}
}
