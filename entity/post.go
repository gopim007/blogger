package entity

type Post struct {
	ID        string `json:"id" gorm:"id"`
	Title     string `json:"title" gorm:"title"`
	Content   string `json:"content" gorm:"content"`
	CreatedAt int    `json:"created_at" gorm:"created_at"`
	UpdatedAt int    `json:"updated_at" gorm:"updated_at"`
}
