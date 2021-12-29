package service

import (
	"github.com/Jangwooo/2022Hackathon/interner/domain/object/request"
	"github.com/Jangwooo/2022Hackathon/interner/mysql"
	"github.com/Jangwooo/2022Hackathon/interner/mysql/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(req request.SignUp) error {
	db := mysql.Connection()

	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	err = db.Create(&model.User{
		ID:          req.ID,
		Password:    string(pwd),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}).Error

	return err
}
