// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	recipes "alta-cookit-be/features/recipes"

	mock "github.com/stretchr/testify/mock"
)

// RecipeData_ is an autogenerated mock type for the RecipeData_ type
type RecipeData_ struct {
	mock.Mock
}

// ActionValidator provides a mock function with given fields: recipeId, userId
func (_m *RecipeData_) ActionValidator(recipeId uint, userId uint) bool {
	ret := _m.Called(recipeId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint, uint) bool); ok {
		r0 = rf(recipeId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteRecipeById provides a mock function with given fields: entity
func (_m *RecipeData_) DeleteRecipeById(entity *recipes.RecipeEntity) error {
	ret := _m.Called(entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) error); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertRecipe provides a mock function with given fields: entity
func (_m *RecipeData_) InsertRecipe(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	ret := _m.Called(entity)

	var r0 *recipes.RecipeEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) (*recipes.RecipeEntity, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) *recipes.RecipeEntity); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recipes.RecipeEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*recipes.RecipeEntity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectRecipeDetailById provides a mock function with given fields: entity
func (_m *RecipeData_) SelectRecipeDetailById(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	ret := _m.Called(entity)

	var r0 *recipes.RecipeEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) (*recipes.RecipeEntity, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) *recipes.RecipeEntity); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recipes.RecipeEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*recipes.RecipeEntity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectRecipes provides a mock function with given fields: entity
func (_m *RecipeData_) SelectRecipes(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	ret := _m.Called(entity)

	var r0 *[]recipes.RecipeEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) (*[]recipes.RecipeEntity, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) *[]recipes.RecipeEntity); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]recipes.RecipeEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*recipes.RecipeEntity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectRecipesTimeline provides a mock function with given fields: entity
func (_m *RecipeData_) SelectRecipesTimeline(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	ret := _m.Called(entity)

	var r0 *[]recipes.RecipeEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) (*[]recipes.RecipeEntity, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) *[]recipes.RecipeEntity); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]recipes.RecipeEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*recipes.RecipeEntity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectRecipesTrending provides a mock function with given fields: entity
func (_m *RecipeData_) SelectRecipesTrending(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	ret := _m.Called(entity)

	var r0 *[]recipes.RecipeEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) (*[]recipes.RecipeEntity, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) *[]recipes.RecipeEntity); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]recipes.RecipeEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*recipes.RecipeEntity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRecipeById provides a mock function with given fields: entity
func (_m *RecipeData_) UpdateRecipeById(entity *recipes.RecipeEntity) error {
	ret := _m.Called(entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*recipes.RecipeEntity) error); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRecipeData_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewRecipeData_ creates a new instance of RecipeData_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRecipeData_(t mockConstructorTestingTNewRecipeData_) *RecipeData_ {
	mock := &RecipeData_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}