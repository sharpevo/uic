package models

type LoginForm struct {
	Email    string `json:"email" form:"email" valid:"Email"`
	Password string `json:"password" form:"password" valid:"Required"`
}
