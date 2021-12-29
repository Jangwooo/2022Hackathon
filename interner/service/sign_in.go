package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object/request"
	"github.com/Jangwooo/2022Hackathon/interner/domain/object/response"
	"github.com/Jangwooo/2022Hackathon/interner/mysql"
	"github.com/Jangwooo/2022Hackathon/interner/mysql/model"
	r "github.com/Jangwooo/2022Hackathon/interner/redis"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrNoMatchingUsers  = fmt.Errorf("no matching users")
	ErrPasswordMismatch = fmt.Errorf("password mismatch")
)

func SignIn(req request.SingIn) (response.SingIn, error) {
	u := model.User{}
	res := response.SingIn{}

	db := mysql.Connection()
	redisDB := r.Connection()
	defer func(r *redis.Client) {
		_ = r.Close()
	}(redisDB)

	err := db.Take(&u, "id = ?", req.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, ErrNoMatchingUsers
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return res, ErrPasswordMismatch
	}

	res.GenerateToken()
	res.Massage = "success"

	redisDB.Set(context.Background(), res.Token, req.ID, 0)
	return res, nil
}
