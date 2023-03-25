package ingredient_details

import "github.com/labstack/echo/v4"

type IngredientDetailEntity struct {
	ID           uint
	UserID       uint `validate:"required"`
	RecipeID     uint `validate:"required"`
	IngredientID uint
	Name         string `validate:"required"`
	Quantity     int
	Unit         string `validate:"required"`
}

type IngredientDetailRequest struct {
	ID           uint `json:"id"`
	UserID       uint
	RecipeID     uint
	IngredientID uint
	Name         string `json:"name" form:"name"`
	Quantity     int    `json:"quantity" form:"quantity"`
	Unit         string `json:"unit" form:"unit"`
}

type IngredientDetailResponse struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Unit     string `json:"unit,omitempty"`
}

type IngredientDetailDelivery_ interface {
	InsertIngredientDetail(e echo.Context) error
	UpdateIngredientDetailById(e echo.Context) error
	DeleteIngredientDetailById(e echo.Context) error
}

type IngredientDetailService_ interface {
	InsertIngredientDetail(ingredientDetailEntity *IngredientDetailEntity) (*IngredientDetailEntity, error)
	UpdateIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
	DeleteIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
}

type IngredientDetailData_ interface {
	InsertIngredientDetail(ingredientDetailEntity *IngredientDetailEntity) (*IngredientDetailEntity, error)
	UpdateIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
	DeleteIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
}
