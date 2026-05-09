package inputs

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
	Position int `json:"position"`
}