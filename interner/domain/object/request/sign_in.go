package request

type SingIn struct {
	ID       string `json:"id,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
