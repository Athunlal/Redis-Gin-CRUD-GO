package domain

import (
	"encoding/json"
)

type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func (u *User) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return data, nil
}
