package model

import "github.com/Jangwooo/2022Hackathon/interner/domain/object"

type Category struct {
	ID   uint   `gorm:"primaryKey; autoIncrementIncrement"`
	Name string `gorm:"not null"`
}

func (Category) TableName() string {
	return "category"
}

func (c Category) ConvertToDTO() object.Category {
	return object.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}
