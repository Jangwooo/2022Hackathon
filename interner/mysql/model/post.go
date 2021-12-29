package model

import (
	"time"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object"
)

type Post struct {
	ID         uint `gorm:"primaryKey; autoIncrementIncrement"`
	UserID     uint
	CategoryID uint
	Title      string    `gorm:"not null"`
	Content    string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`

	Category Category
}

func (Post) TableName() string {
	return "post"
}

func (p Post) ConvertToDTO() object.Post {
	return object.Post{
		ID:     p.ID,
		UserID: p.UserID,
		Category: object.Category{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
	}
}
