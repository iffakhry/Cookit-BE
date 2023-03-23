package steps

import "github.com/labstack/echo/v4"

type StepEntity struct {
	ID       uint
	UserID   uint
	UserRole string
	RecipeID uint
	Name     string
}

type StepRequest struct {
	ID       uint    
	UserID   uint    
	UserRole string  
	RecipeID uint    
	Name     string  `json:"name"`
}

type StepResponse struct {
	ID    uint    `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
}

type StepDelivery_ interface {
	InsertStep(e echo.Context) error
	UpdateStepById(e echo.Context) error
	DeleteStepById(e echo.Context) error
}

type StepService_ interface {
	InsertStep(stepEntity *StepEntity) (*StepEntity, error)
	UpdateStepById(stepEntity *StepEntity) error
	DeleteStepById(stepEntity *StepEntity) error
}

type StepData_ interface {
	InsertStep(stepEntity *StepEntity) (*StepEntity, error)
	UpdateStepById(istepEntity *StepEntity) error
	DeleteStepById(stepEntity *StepEntity) error
}
