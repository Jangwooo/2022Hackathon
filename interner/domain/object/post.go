package object

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	Category  Category  `json:"category"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
