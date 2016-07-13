package models

type RegisterForm struct {
	Email           string `form:"email" valid:"Required;Email"`
	Name            string `form:"name" valid:"Required"`
	Password        string `form:"password" valid:"MinSize(6);MaxSize(15)"`
	ConfirmPassword string `form:"password_confirmation" valid:"MinSize(6);MaxSize(15)"`
}
