package object

import (
	"time"
)

type Post struct {
	ID         uint      `json:"id,omitempty"`
	UserID     string    `json:"user_id,omitempty"`
	CategoryID uint      `json:"category_id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	Detail     string    `json:"detail,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Icon       string    `json:"icon,omitempty"`
}
