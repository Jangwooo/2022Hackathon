package object

type User struct {
	ID          string `json:"id,omitempty"`
	Password    string `json:"password,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
