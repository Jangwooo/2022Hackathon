package model

import (
	"time"
)

type Post struct {
	ID         uint `gorm:"primaryKey; autoIncrementIncrement"`
	UserID     string
	CategoryID uint
	Title      string    `gorm:"not null"`
	Content    string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`

	Category Category
}

func (Post) TableName() string {
	return "post"
}
