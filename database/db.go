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

// NewConnection - A function to set up databse connection with a connecton string
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

// CloseConnection - closing database connection
func (db *DB) CloseConnection() {
	sqlConn, _ := db.Conn.DB()
	sqlConn.Close()
}

// CreatePost - create post DB repo
func (db *DB) CreatePost(post entity.Post) (*entity.Post, error) {
	var createdPost *entity.Post
	result := db.Conn.Raw(`INSERT INTO posts (id,title,content,created_at,updated_at) VALUES(?,?,?,?,?) retuning *`, post.ID, post.Title, post.Content, post.CreatedAt, post.UpdatedAt).Scan(&createdPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			return nil, errors.New(constants.DUPLICATE_ERROR)
		}
		return nil, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return createdPost, nil

}

// GetPosts  - Get posts DB repo
func (db *DB) GetPosts(postId string) ([]entity.Post, error) {
	posts := make([]entity.Post, 0)
	var result *gorm.DB
	if postId == "" {
		result = db.Conn.Raw(`select * from posts`).Scan(&posts)
	} else {
		result = db.Conn.Raw(`select * from posts where id = ?`, postId).Scan(&posts)
	}

	if result.Error != nil {
		return nil, errors.New(constants.SOMETHING_WENT_WRONG)
	}
	return posts, nil
}

// UpdatePostByID - update post by id db repo
func (db *DB) UpdatePostByID(post entity.Post) (*entity.Post, error) {
	var reultPost entity.Post
	result := db.Conn.Raw(`update posts set title = ?, content = ?, updated_at = ? where id = ? returning *`, post.Title, post.Content, post.UpdatedAt, post.ID).Scan(&reultPost)
	if result.Error != nil {
		return nil, errors.New(constants.SOMETHING_WENT_WRONG)
	}
	return &reultPost, nil
}

// DeletePostByID - delete post by id db repo
func (db *DB) DeletePostByID(postId string) (*entity.Post, error) {
	var deletedPost entity.Post
	result := db.Conn.Raw(`delete from posts where id = ? returning *`, postId).Scan(&deletedPost)
	if result.Error != nil {
		return nil, errors.New(constants.SOMETHING_WENT_WRONG)
	}
	return &deletedPost, nil
}
