package repository

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type PhotoRepository struct{
	db * gorm.DB
}

func NewPhotoRepository(db *gorm.DB)(*PhotoRepository){
	return &PhotoRepository{db}
}

func (pr *PhotoRepository) AddPhoto (addPhoto model.Photo) (model.Photo,error){
	tx:= pr.db.Create(&addPhoto)
	
	return addPhoto, tx.Error
}

func (pr *PhotoRepository) GetPhotoUserId (userId string) ([]model.Photo,error){
	var photo = []model.Photo{}
	tx:= pr.db.Find(&photo,"user_id= ?",userId)
	
	return photo, tx.Error	
}

func (pr *PhotoRepository) GetPhotoById (id string) (model.Photo,error){
	var photo model.Photo
	tx:= pr.db.First(&photo,"id = ?",id)
	
	return photo, tx.Error	
}


func (pr *PhotoRepository) UpdatePhoto (updatePhoto model.Photo, id string) (model.Photo,error){
	tx:= pr.db.Model(&updatePhoto).Where("id = ?",id).Updates(&updatePhoto)

	return updatePhoto, tx.Error	
}




// Is it possible to run it ?
func (pr *PhotoRepository) DeletePhoto (id string) (error){
	var deletePhoto model.Photo
	tx:= pr.db.Where("id = ?",id).Delete(&deletePhoto)

	return tx.Error	
}


