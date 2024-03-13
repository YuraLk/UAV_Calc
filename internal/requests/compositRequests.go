package requests

type CreateComposit struct {
	Name string `json:"name" binding:"required"`
}

type UpdateComposit struct {
	Id   uint   `json:"ID" binding:"required"`
	Name string `json:"name" binding:"required"`
}
