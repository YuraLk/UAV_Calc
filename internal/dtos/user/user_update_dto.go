package user

type UserUpdateDTO struct {
	Name   string `json:"name" binding:"required,fullName"`
	Email  string `json:"email" binding:"required,email"`
	Phone  string `json:"phone" binding:"required"`
	Device string `json:"device" binding:"required"`
}
