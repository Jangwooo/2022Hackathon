package response

import "github.com/Jangwooo/2022Hackathon/interner/domain/object"

type Posts struct {
	Massage string     `json:"massage"`
	Posts   []PostItem `json:"posts"`
}

type Post struct {
	Massage string      `json:"massage"`
	User    object.User `json:"user"`
	Post    object.Post `json:"post"`
}

type PostItem struct {
	User object.User `json:"user"`
	Post object.Post `json:"post"`
}
