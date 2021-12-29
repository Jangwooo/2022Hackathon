package object

import (
	"time"

	"github.com/Jangwooo/2022Hackathon/interner/mysql/model"
)

type Post struct {
	ID        uint      `json:"id,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	Category  Category  `json:"category"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (p Post) ConvertToDAO() model.Post {
	return model.Post{
		ID:         p.ID,
		UserID:     p.UserID,
		CategoryID: p.Category.ID,
		Title:      p.Title,
		Content:    p.Content,
		CreatedAt:  p.CreatedAt,
		Category: model.Category{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
	}
}
