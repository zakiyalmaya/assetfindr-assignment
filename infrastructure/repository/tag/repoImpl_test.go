package tag

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDBMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}

	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatalf("An error occurred while creating gorm db mock: %s", err)
	}

	return sqldb, gormdb, mock
}

func TestGetOrCreate(t *testing.T) {
	sqldb, gormdb, mock := SetupDBMock(t)
	defer sqldb.Close()

	tag := &model.Tag{
		Label: "Tag1",
	}

	testCases := []struct {
		name    string
		request *model.Tag
		mock    func()
		tx      []*gorm.DB
		wantErr bool
	}{
		{
			name:    "Given valid request when create or get tag from db then return success",
			request: tag,
			mock: func() {
				mock.ExpectQuery("SELECT * FROM \"tags\" WHERE \"tags\".\"label\" = $1 ORDER BY \"tags\".\"id\" LIMIT $2").
					WithArgs(tag.Label, 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			tx:      []*gorm.DB{gormdb.Model(tag)},
			wantErr: false,
		},
		{
			name:    "Given error when create or get tag from db then return error",
			request: tag,
			mock: func() {
				mock.ExpectQuery("SELECT * FROM \"tags\" WHERE \"tags\".\"label\" = $1 ORDER BY \"tags\".\"id\" LIMIT $2").
					WithArgs(tag.Label, 1).WillReturnError(errors.New("error"))
			},
			tx:      []*gorm.DB{gormdb.Model(tag)},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewTagRepository(gormdb)

			tc.mock()
			_, err := repo.GetOrCreate(tc.request, tc.tx[0])
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}
