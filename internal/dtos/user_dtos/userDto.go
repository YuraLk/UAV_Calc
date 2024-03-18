package user_dtos

type UserDTO struct {
	Id    uint   `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"access"`
}
