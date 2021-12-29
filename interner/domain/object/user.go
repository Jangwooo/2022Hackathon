package object

import "github.com/Jangwooo/2022Hackathon/interner/mysql/model"

type User struct {
	ID          string `json:"id,omitempty"`
	Password    string `json:"password,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func (u User) ConvertToDAO() model.User {
	return model.User{
		ID:          u.Name,
		Password:    u.Password,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}
}
