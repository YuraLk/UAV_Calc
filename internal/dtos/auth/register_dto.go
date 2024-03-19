package auth

type RegisterDTO struct {
	Name     string `json:"name" binding:"required,fullName"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=10,max=40,safety"`
	Device   string `json:"device" binding:"required"`
}
