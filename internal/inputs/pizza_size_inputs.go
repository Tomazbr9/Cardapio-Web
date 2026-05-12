package inputs

type CreatePizzaSizeInput struct {
	Name string `json:"name" binding:"required"`
	Slices int `json:"slices" binding:"required"`
	MaxFlavors int `json:"max_flavors" binding:"required"`
}