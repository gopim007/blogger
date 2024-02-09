package entity

import (
	"blogger/constants"
	"errors"
)

type Post struct {
	ID        string `json:"id" gorm:"id"`
	Title     string `json:"title" gorm:"title"`
	Content   string `json:"content" gorm:"content"`
	CreatedAt int    `json:"created_at" gorm:"created_at"`
	UpdatedAt int    `json:"updated_at" gorm:"updated_at"`
}

func (post Post) Validate() error {
	if post.ID == "" || post.Title == "" || post.Content == "" || post.CreatedAt == 0 {
		return errors.New(constants.INVALID_DATA)
	}
	return nil
}
