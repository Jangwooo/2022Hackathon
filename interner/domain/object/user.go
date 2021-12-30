package object

type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
