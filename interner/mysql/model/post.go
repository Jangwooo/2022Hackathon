package model

import (
	"time"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object"
)

type Post struct {
	ID         uint `gorm:"primaryKey; autoIncrementIncrement"`
	UserID     string
	CategoryID uint
	Title      string    `gorm:"not null"`
	Content    string    `gorm:"not null"`
	Detail     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`

	Category Category
}

func (Post) TableName() string {
	return "post"
}

func (p Post) ToObject() object.Post {
	return object.Post{
		ID:         p.ID,
		UserID:     p.UserID,
		Title:      p.Title,
		CategoryID: p.CategoryID,
		Content:    p.Content,
		Detail:     p.Detail,
		CreatedAt:  p.CreatedAt,
	}
}
