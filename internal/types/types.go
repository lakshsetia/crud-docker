package types

import "fmt"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (user *User) Validate() error {
	if user.Name == "" {
		return fmt.Errorf("missing field: name")
	}
	if user.Email == "" {
		return fmt.Errorf("missing field: email")
	}
	if user.Age < 0 {
		return fmt.Errorf("invalid field: age")
	}
	return nil
}