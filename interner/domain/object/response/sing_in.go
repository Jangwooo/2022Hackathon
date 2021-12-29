package response

import (
	"crypto/rand"
	"fmt"
)

type SingIn struct {
	Massage string `json:"massage"`
	Token   string `json:"token,omitempty"`
}

func (t *SingIn) GenerateToken() {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	t.Token = fmt.Sprintf("%x", b)[:8]
}
