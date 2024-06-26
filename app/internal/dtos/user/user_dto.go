package user

type UserDTO struct {
	Id         uint   `json:"ID"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Role       string `json:"access"`
}
