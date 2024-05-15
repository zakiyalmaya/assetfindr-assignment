package post

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

func TestCreate(t *testing.T) {
	sqldb, gormdb, mock := SetupDBMock(t)
	defer sqldb.Close()

	tags := []*model.Tag{
		{Label: "Tag1"},
		{Label: "Tag2"},
	}

	req := &model.Post{
		Title:   "title",
		Content: "content",
		Tags:    tags,
	}

	testCases := []struct {
		name    string
		request *model.Post
		mock    func()
		tx      []*gorm.DB
		wantErr bool
	}{
		{
			name:    "Given valid request when create post to db then return success",
			request: req,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery("INSERT INTO \"posts\" (\"title\",\"content\") VALUES ($1,$2) RETURNING \"id\"").
					WithArgs(req.Title, req.Content).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectQuery("INSERT INTO \"tags\" (\"label\") VALUES ($1),($2) ON CONFLICT DO NOTHING RETURNING \"id\"").
					WithArgs(req.Tags[0].Label, req.Tags[1].Label).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

				mock.ExpectExec("INSERT INTO \"post_tags\" (\"post_id\",\"tag_id\") VALUES ($1,$2),($3,$4) ON CONFLICT DO NOTHING").
					WithArgs(1, 1, 1, 2).WillReturnResult(sqlmock.NewResult(0, 2))

				mock.ExpectCommit()
			},
			tx:      []*gorm.DB{},
			wantErr: false,
		},
		{
			name:    "Given error when create post to db then return error",
			request: req,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery("INSERT INTO \"posts\" (\"title\",\"content\",\"id\") VALUES ($1,$2,$3) RETURNING \"id\"").
					WithArgs(req.Title, req.Content, req.ID).
					WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			tx:      []*gorm.DB{gormdb.Begin()},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewPostRepository(gormdb)

			tc.mock()
			err := repo.Create(tc.request, tc.tx...)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	sqldb, gormdb, mock := SetupDBMock(t)
	defer sqldb.Close()

	testCases := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when get all post then return success",
			mock: func() {
				mock.ExpectQuery("SELECT * FROM \"posts\"").
					WithoutArgs().
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content"}).
						AddRow(1, "Title", "Content"))

				mock.ExpectQuery("SELECT * FROM \"post_tags\" WHERE \"post_tags\".\"post_id\" = $1").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"post_id", "tag_id"}).
						AddRow(1, 1))

				mock.ExpectQuery("SELECT * FROM \"tags\" WHERE \"tags\".\"id\" = $1").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "label"}).
						AddRow(1, "Label"))
			},
			wantErr: false,
		},
		{
			name: "Given error when get all post then return error",
			mock: func() {
				mock.ExpectQuery("SELECT * FROM \"posts\"").
					WithoutArgs().WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewPostRepository(gormdb)

			tc.mock()
			_, err := repo.GetAll()
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	sqldb, gormdb, mock := SetupDBMock(t)
	defer sqldb.Close()

	testCases := []struct {
		name    string
		request int
		mock    func()
		tx      []*gorm.DB
		wantErr bool
	}{
		{
			name:    "Given valid request when get by id then return success",
			request: 1,
			mock: func() {
				mock.ExpectQuery("SELECT * FROM \"posts\" WHERE \"posts\".\"id\" = $1 ORDER BY \"posts\".\"id\" LIMIT $2").
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content"}).
						AddRow(1, "Title", "Content"))

				mock.ExpectQuery("SELECT * FROM \"post_tags\" WHERE \"post_tags\".\"post_id\" = $1").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"post_id", "tag_id"}).
						AddRow(1, 1))

				mock.ExpectQuery("SELECT * FROM \"tags\" WHERE \"tags\".\"id\" = $1").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "label"}).
						AddRow(1, "Label"))
			},
			tx:      []*gorm.DB{},
			wantErr: false,
		},
		{
			name:    "Given valid request when get by id then return success",
			request: 1,
			mock: func() {
				mock.ExpectQuery("SELECT * FROM \"posts\" WHERE \"posts\".\"id\" = $1 ORDER BY \"posts\".\"id\" LIMIT $2").
					WithArgs(1, 1).
					WillReturnError(errors.New("error"))
			},
			tx:      []*gorm.DB{gormdb.Begin()},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewPostRepository(gormdb)

			tc.mock()
			_, err := repo.GetByID(tc.request, tc.tx...)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	sqldb, gormdb, mock := SetupDBMock(t)
	defer sqldb.Close()

	tags := []*model.Tag{
		{Label: "Tag1"},
		{Label: "Tag2"},
	}

	req := &model.Post{
		ID:      1,
		Title:   "title",
		Content: "content",
		Tags:    tags,
	}

	testCases := []struct {
		name    string
		request *model.Post
		mock    func()
		tx      []*gorm.DB
		wantErr bool
	}{
		{
			name:    "Given valid request when update post to db then return success",
			request: req,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("UPDATE \"posts\" SET \"title\"=$1,\"content\"=$2 WHERE \"id\" = $3").
					WithArgs(req.Title, req.Content, req.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery("INSERT INTO \"tags\" (\"label\") VALUES ($1),($2) ON CONFLICT DO NOTHING RETURNING \"id\"").
					WithArgs(req.Tags[0].Label, req.Tags[1].Label).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

				mock.ExpectExec("INSERT INTO \"post_tags\" (\"post_id\",\"tag_id\") VALUES ($1,$2),($3,$4) ON CONFLICT DO NOTHING").
					WithArgs(1, 1, 1, 2).WillReturnResult(sqlmock.NewResult(0, 2))

				mock.ExpectCommit()
			},
			tx:      []*gorm.DB{},
			wantErr: false,
		},
		{
			name:    "Given error when update post to db then return error",
			request: req,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("UPDATE \"posts\" SET \"title\"=$1,\"content\"=$2 WHERE \"id\" = $3").
					WithArgs(req.Title, req.Content, req.ID).
					WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			tx:      []*gorm.DB{gormdb.Begin()},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewPostRepository(gormdb)

			tc.mock()
			err := repo.Update(tc.request, tc.tx...)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestDelete(t *testing.T) {
	sqldb, gormdb, mock := SetupDBMock(t)
	defer sqldb.Close()

	testCases := []struct {
		name    string
		request int
		mock    func()
		tx      []*gorm.DB
		wantErr bool
	}{
		{
			name:    "Given valid request when delete by id then return success",
			request: 1,
			mock: func() {
				mock.ExpectExec("DELETE FROM post_tags WHERE post_id = $1").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM \"posts\" WHERE \"posts\".\"id\" = $1").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			tx:      []*gorm.DB{},
			wantErr: false,
		},
		{
			name:    "Given error delete post_tags when delete by id then return error",
			request: 1,
			mock: func() {
				mock.ExpectExec("DELETE FROM post_tags WHERE post_id = $1").
					WithArgs(1).WillReturnError(errors.New("error"))
			},
			tx:      []*gorm.DB{gormdb.Begin()},
			wantErr: true,
		},
		{
			name:    "Given error delete posts when delete by id then return error",
			request: 1,
			mock: func() {
				mock.ExpectExec("DELETE FROM post_tags WHERE post_id = $1").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM \"posts\" WHERE \"posts\".\"id\" = $1").
					WithArgs(1).WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			tx:      []*gorm.DB{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewPostRepository(gormdb)

			tc.mock()
			err := repo.Delete(tc.request, tc.tx...)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestAssosiate(t *testing.T) {
	sqldb, gormdb, mock := SetupDBMock(t)
	defer sqldb.Close()

	tags := []*model.Tag{
		{Label: "Tag1"},
		{Label: "Tag2"},
	}

	post := &model.Post{
		ID:      1,
		Title:   "title",
		Content: "content",
		Tags:    tags,
	}

	type request struct {
		post *model.Post
		tags []*model.Tag
	}

	testCases := []struct {
		name    string
		request request
		mock    func()
		tx      []*gorm.DB
		wantErr bool
	}{
		{
			name: "Given valid request when assosiate post then return success",
			request: request{
				post: post,
				tags: tags,
			},
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery("INSERT INTO \"tags\" (\"label\") VALUES ($1),($2) ON CONFLICT DO NOTHING RETURNING \"id\"").
					WithArgs(tags[0].Label, tags[1].Label).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

				mock.ExpectExec("INSERT INTO \"post_tags\" (\"post_id\",\"tag_id\") VALUES ($1,$2),($3,$4) ON CONFLICT DO NOTHING").
					WithArgs(1, 1, 1, 2).WillReturnResult(sqlmock.NewResult(0, 2))

				mock.ExpectCommit()

				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM  \"post_tags\" WHERE \"post_tags\".\"post_id\" = $1 AND \"post_tags\".\"tag_id\" NOT IN ($2,$3)").
					WithArgs(1, 1, 2).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			tx:      []*gorm.DB{},
			wantErr: false,
		},
		{
			name: "Given error when assosiate post then return error",
			request: request{
				post: post,
				tags: tags,
			},
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery("INSERT INTO \"tags\" (\"label\") VALUES ($1),($2) ON CONFLICT DO NOTHING RETURNING \"id\"").
					WithArgs(tags[0].Label, tags[1].Label).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

				mock.ExpectExec("INSERT INTO \"post_tags\" (\"post_id\",\"tag_id\") VALUES ($1,$2),($3,$4) ON CONFLICT DO NOTHING").
					WithArgs(1, 1, 1, 2).WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			tx:      []*gorm.DB{gormdb.Begin()},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewPostRepository(gormdb)

			tc.mock()
			err := repo.Assosiate(tc.request.post, tc.request.tags, tc.tx...)
			if err != nil && !tc.wantErr {
				t.Fatalf(err.Error())
			}
		})
	}
}
