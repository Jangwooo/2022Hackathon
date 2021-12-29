package service

import (
	"github.com/Jangwooo/2022Hackathon/interner/domain/object/request"
	"github.com/Jangwooo/2022Hackathon/interner/mysql"
	"github.com/Jangwooo/2022Hackathon/interner/mysql/model"
)

func CreatePost(req request.CreatePost, userID string) error {
	db := mysql.Connection()

	return db.Create(&model.Post{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Content:    req.Content,
	}).Error
}
