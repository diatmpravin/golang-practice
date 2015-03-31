package main

import (
	"encoding/json"
	"fmt"
	// "os"
)

type omit bool

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func NewUser() (*User, error) {
// 	u := &User{}
// 	return u, json.Unmarshal([]byte(`{
// 	  "email": "test@example.com",
// 	  "password": "secret"
// 	}`), u)
// }

func main() {
	// u, _ := NewUser()

	u := &User{}
	b := []byte(`{"email": "test@example.com","password": "secret"}`)

	err := json.Unmarshal(b, u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", u)

	// Omit the password:
	// json.NewEncoder(os.Stdout).Encode(struct {
	// 	*User
	// 	Password omit `json:"password,omitempty"`
	// }{
	// 	User: u,
	// })

	// type PublicUser struct {
	// 	*User
	// 	Password omit `json:"password,omitempty"`
	// }

	// pu, _ := json.Marshal(PublicUser{User: u})
	// fmt.Printf(pu)
}
