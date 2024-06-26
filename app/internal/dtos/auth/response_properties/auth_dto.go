package response_properties

type AuthDTO struct {
	Id         uint    `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic string  `json:"patronymic"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Role       *string `json:"access"`
}
