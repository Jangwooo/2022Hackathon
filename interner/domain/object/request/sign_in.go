package request

type SingIn struct {
	ID          string `json:"id,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
	Name        string `json:"name,omitempty" binding:"required"`
	PhoneNumber string `json:"phone_number,omitempty" binding:"required"`
}
