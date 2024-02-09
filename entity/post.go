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

// Validate - method to validate the entity
func (post Post) Validate(scope string) error {
	switch scope {
	default:
	case constants.CREATE:
		if post.ID == "" || post.Title == "" || post.Content == "" || post.CreatedAt == 0 {
			return errors.New(constants.INVALID_DATA)
		}
	case constants.UPDATE:
		if post.ID == "" || post.Title == "" || post.Content == "" || post.UpdatedAt == 0 {
			return errors.New(constants.INVALID_DATA)
		}
	}

	return nil
}
