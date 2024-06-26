package auth

type RegisterDTO struct {
	Name       string `json:"name" binding:"required,min=3,max=64"`
	Surname    string `json:"surname" binding:"required,min=3,max=64"`
	Patronymic string `json:"patronymic" binding:"max=64"`
	Email      string `json:"email" binding:"required,email,max=64"`
	Phone      string `json:"phone" binding:"required,min=11,max=32"`
	Password   string `json:"password" binding:"required,min=10,max=40,isSafety"`
	Device     string `json:"device" binding:"required"`
}
