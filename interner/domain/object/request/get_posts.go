package request

type GetPosts struct {
	Category int `form:"category"`
}
