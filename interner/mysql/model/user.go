package model

import "github.com/Jangwooo/2022Hackathon/interner/domain/object"

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

func (u User) ToObject() object.User {
	return object.User{
		ID:          u.ID,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}
}
