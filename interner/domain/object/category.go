package object

import "github.com/Jangwooo/2022Hackathon/interner/mysql/model"

type Category struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c Category) ConvertToDAO() model.Category {
	return model.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}
