package request_properties

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Device   string `json:"device" binding:"required"`
}
