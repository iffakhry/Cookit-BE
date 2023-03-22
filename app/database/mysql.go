package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"alta-cookit-be/app/config"
	_cartModel "alta-cookit-be/features/carts/models"
	_commentModel "alta-cookit-be/features/comments/models"
	_followerModel "alta-cookit-be/features/followers/models"
	_imageModel "alta-cookit-be/features/images/models"
	_ingredientDetailModel "alta-cookit-be/features/ingredient_details/models"
	_ingredientModel "alta-cookit-be/features/ingredients/models"
	_likeModel "alta-cookit-be/features/likes/models"
	_recipeModel "alta-cookit-be/features/recipes/models"
	_stepModel "alta-cookit-be/features/steps/models"
	_transactionDetailModel "alta-cookit-be/features/transaction_details/models"
	_transactionModel "alta-cookit-be/features/transactions/models"
	_userModel "alta-cookit-be/features/users/models"
	
)

func InitDB(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to DB", err.Error())
		return nil
	}

	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(
		_userModel.User{}, _cartModel.Cart{}, _commentModel.Comment{},
		_followerModel.Follower{}, _imageModel.Image{}, _ingredientDetailModel.IngredientDetail{},
		_ingredientModel.Ingredient{},_likeModel.Like{}, _recipeModel.Recipe{}, _stepModel.Step{}, 
		_transactionModel.Transaction{}, _transactionDetailModel.TransactionsDetail{},
	)
}