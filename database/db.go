package database

import (
	"blogger/constants"
	"blogger/entity"
	"errors"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Conn *gorm.DB
}

func NewConnection(connectionString string) (*DB, error) {
	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	return &DB{
		Conn: conn,
	}, nil
}

func (db *DB) CreatePost(post entity.Post) (*entity.Post, error) {
	var createdPost *entity.Post
	result := db.Conn.Raw(`INSERT INTO posts (id,title,content,created_at,updated_at) VALUES(?,?,?,?,?)`, post.ID, post.Title, post.Content, post.CreatedAt, post.UpdatedAt).Scan(&createdPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			return nil, errors.New(constants.DUPLICATE_ERROR)
		}
		return nil, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return createdPost, nil

}
