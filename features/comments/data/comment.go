package data

import (
	"alta-cookit-be/app/storage"
	"alta-cookit-be/features/comments"
	_commentModel "alta-cookit-be/features/comments/models"
	"alta-cookit-be/features/users"
	_userModel "alta-cookit-be/features/users/data"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CommentData struct {
	db       *gorm.DB
	userData users.UserData
}

func New(db *gorm.DB, userData users.UserData) comments.CommentData_ {
	return &CommentData{
		db:       db,
		userData: userData,
	}
}

func (d *CommentData) ActionValidator(recipeId, userId uint) bool {
	tempGorm := _commentModel.Comment{}
	d.db.Model(&tempGorm).Where("recipe_id = ? AND user_id = ?", recipeId, userId).Find(&tempGorm)

	return tempGorm.ID != 0
}

func (d *CommentData) SelectCommentById(id uint) *_commentModel.Comment {
	tempGorm := _commentModel.Comment{}
	d.db.Where("id = ?", id).Find(&tempGorm)

	if tempGorm.ID == 0 {
		return nil
	}
	return &tempGorm
}

func (d *CommentData) SelectCommentsByRecipeId(entity *comments.CommentEntity) (*[]comments.CommentEntity, error) {
	gorms := []_commentModel.Comment{}

	tx := d.db.Where("recipe_id = ?", entity.RecipeID).Limit(entity.DataLimit).Offset(entity.DataOffset).Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []comments.CommentEntity{}
	for _, comment := range gorms {
		userEntity := d.userData.SelectUserById(users.Core{ID: comment.UserID})
		userModel := _userModel.CoreToModel(*userEntity)
		entities = append(entities, *ConvertToEntity(&comment, &userModel))
	}

	return &entities, nil
}

func (d *CommentData) InsertComment(entity *comments.CommentEntity) (*comments.CommentEntity, error) {
	gorm := ConvertToGorm(entity)

	if entity.Image != nil {
		// Local
		// urlImage, err := storage.UploadFile(entity.Image, entity.ImageName)

		// Google Cloud Storage
		urlImage, err := storage.GetStorageClient().UploadFile(entity.Image, entity.ImageName)

		if err != nil {
			return nil, err
		}
		gorm.UrlImage = urlImage
	}

	tx := d.db.Create(gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		if strings.Contains(tx.Error.Error(), "recipe_id") {
			return nil, errors.New(consts.RECIPE_InvalidRecipe)
		}
		return nil, tx.Error
	}
	return ConvertToEntity(gorm), nil
}

func (d *CommentData) UpdateCommentById(entity *comments.CommentEntity) (*comments.CommentEntity, error) {
	gorm, tempGorm := ConvertToGorm(entity), d.SelectCommentById(entity.ID)

	if tempGorm == nil {
		return nil, errors.New(consts.GORM_RecordNotFound)
	}
	if entity.Image != nil {
		if tempGorm.UrlImage != "" {
			// Local
			// err := storage.DeleteFile(tempGorm.UrlImage)

			// Google Cloud Storage
			err := storage.GetStorageClient().DeleteFile(tempGorm.UrlImage)

			if err != nil {
				return nil, err
			}
		}

		// Local
		// urlImage, err := storage.UploadFile(entity.Image, entity.ImageName)

		// Google Cloud Storage
		urlImage, err := storage.GetStorageClient().UploadFile(entity.Image, entity.ImageName)

		if err != nil {
			return nil, err
		}
		gorm.UrlImage = urlImage
	}

	tx := d.db.Where("id = ?", entity.ID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		if strings.Contains(tx.Error.Error(), "recipe_id") {
			return nil, errors.New(consts.RECIPE_InvalidRecipe)
		}
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New(consts.GORM_RecordNotFound)
	}
	return ConvertToEntity(gorm), nil
}

func (d *CommentData) DeleteCommentById(entity *comments.CommentEntity) error {
	tx := d.db.Unscoped().Where("id = ?", entity.ID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}
