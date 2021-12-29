package response

import "github.com/Jangwooo/2022Hackathon/interner/domain/object"

type Post struct {
	Massage string      `json:"massage"`
	User    object.User `json:"user,omitempty"`
	Post    object.Post `json:"post,omitempty"`
}
