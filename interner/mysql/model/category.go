package model

type Category struct {
	ID   uint   `gorm:"primaryKey; autoIncrementIncrement"`
	Name string `gorm:"not null"`
}

func (Category) TableName() string {
	return "category"
}
