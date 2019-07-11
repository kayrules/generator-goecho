package model

// User struct
type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}
