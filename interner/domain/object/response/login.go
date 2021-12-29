package response

type User struct {
	Massage string `json:"massage"`
	Token   string `json:"token,omitempty"`
}
