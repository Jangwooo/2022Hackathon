package model

type User struct {
	ID          string `gorm:"primaryKey"`
	Password    string `gorm:"not null"`
	Name        string `gorm:"not null"`
	PhoneNumber string `gorm:"not null"`

	Posts []Post
}

func (User) TableName() string {
	return "user"
}
