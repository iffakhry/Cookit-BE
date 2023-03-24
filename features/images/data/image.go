package data

import (
	"alta-cookit-be/app/storage"
	"alta-cookit-be/features/images"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type ImageData struct {
	db *gorm.DB
}

func New(db *gorm.DB) images.ImageData_ {
	return &ImageData{
		db: db,
	}
}

func (d *ImageData) InsertImage(entities *[]images.ImageEntity) (*[]images.ImageEntity, error) {
	gorms := *ConvertToGorms(entities)

	for index, entity := range *entities {
		if entity.Image != nil {
			urlImage, err := storage.GetStorageClient().UploadFile(entity.Image, entity.ImageName)
			if err != nil {
				return nil, err
			}
			gorms[index].UrlImage = urlImage
		}
	}

	if len(*entities) != 0 {
		tx := d.db.Create(&gorms)
		if tx.Error != nil {
			if strings.Contains(tx.Error.Error(), "recipe_id") {
				return nil, errors.New(consts.RECIPE_InvalidRecipe)
			}
			return nil, tx.Error
		}
	}
	return ConvertToEntities(&gorms), nil
}

func (d *ImageData) UpdateImageById(entity *images.ImageEntity) (*images.ImageEntity, error) {
	gorm := ConvertToGorm(entity)

	tx := d.db.Where("id = ? AND recipe_id = ?", entity.ID, entity.RecipeID).Find(&gorm)
	if tx.RowsAffected == 0 {
		return nil, errors.New(consts.GORM_RecordNotFound)
	}

	if entity.Image != nil {
		if gorm.UrlImage != "" {
			err := storage.GetStorageClient().DeleteFile(gorm.UrlImage)
			if err != nil {
				return nil, err
			}
		}

		urlImage, err := storage.GetStorageClient().UploadFile(entity.Image, entity.ImageName)
		if err != nil {
			return nil, err
		}
		gorm.UrlImage = urlImage
	}

	tx = d.db.Where("id = ? AND recipe_id = ?", entity.ID, entity.RecipeID).Updates(&gorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ConvertToEntity(gorm), nil
}

func (d *ImageData) DeleteImageById(entity *images.ImageEntity) error {
	tx := d.db.Where("id = ? AND recipe_id = ?", entity.ID, entity.RecipeID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}