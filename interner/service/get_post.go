package service

import (
	"fmt"
	"math/rand"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object/response"
	"github.com/Jangwooo/2022Hackathon/interner/mysql"
	"github.com/Jangwooo/2022Hackathon/interner/mysql/model"
)

func GetPost(postID string) response.Post {
	db := mysql.Connection()

	res := response.Post{}
	var post model.Post
	var user model.User

	if err := db.Find(&post, "id = ?", postID).Error; err != nil {
		panic(err.Error())
	}

	if err := db.Select("name, phone_number").Find(&user, "id = ?", post.UserID).Error; err != nil {
		panic(err.Error())
	}

	res.Post = post.ToObject()
	res.User = user.ToObject()
	switch post.CategoryID {
	case 1:
		res.Post.Icon = fmt.Sprintf("15.165.88.215:8000/image/select%d.png", rand.Intn(2)+1)
	case 2:
		res.Post.Icon = fmt.Sprintf("15.165.88.215:8000/image/lightnig.png")
	case 3:
		res.Post.Icon = fmt.Sprintf("15.165.88.215:8000/image/persuade.png")
	}

	res.Massage = "success"

	return res
}
