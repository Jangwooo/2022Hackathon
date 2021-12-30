package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object/request"
	"github.com/Jangwooo/2022Hackathon/interner/domain/object/response"
	"github.com/Jangwooo/2022Hackathon/interner/mysql"
	"github.com/Jangwooo/2022Hackathon/interner/mysql/model"
)

func GetPosts(req request.GetPosts) response.Posts {
	db := mysql.Connection()
	rand.Seed(time.Now().UnixNano())

	var posts []model.Post
	var user model.User
	var res response.Posts
	var temp response.PostItem

	res.Posts = []response.PostItem{}

	if req.Category != 0 {
		if err := db.Select("id, user_id, title, category_id, created_at").Find(&posts, "category_id = ?", req.Category).Error; err != nil {
			panic(err.Error())
		}
	} else {
		if err := db.Find(&posts).Error; err != nil {
			panic(err.Error())
		}
	}

	for _, post := range posts {
		if err := db.Select("name, phone_number").Find(&user, "id = ?", post.UserID).Error; err != nil {
			panic(err.Error())
		}
		temp.User = user.ToObject()
		temp.Post = post.ToObject()
		switch post.CategoryID {
		case 1:
			temp.Post.Icon = fmt.Sprintf("15.165.88.215:8000/image/select%d.png", rand.Intn(2)+1)
		case 2:
			temp.Post.Icon = fmt.Sprintf("15.165.88.215:8000/image/lightnig.png")
		case 3:
			temp.Post.Icon = fmt.Sprintf("15.165.88.215:8000/image/persuade.png")
		}

		res.Posts = append(res.Posts, temp)
	}
	res.Massage = "success"

	return res
}
