package repository

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type PhotoRepository struct{
	db * gorm.DB
}


type IPhotoRepository interface{
	AddPhoto(addPhoto model.Photo) (model.Photo,error)
	GetAllPhoto() ([]model.Photo,error)
	GetPhotoByUserId(userId string) ([]model.Photo,error)
	GetPhotoByPhotoId(id string) (model.Photo,error)
	UpdatePhoto(updatePhoto model.Photo, id string) (model.Photo,error)
	DeletePhoto(id string) (error)
}

func NewPhotoRepository(db *gorm.DB)(*PhotoRepository){
	return &PhotoRepository{db}
}

func (pr *PhotoRepository) AddPhoto (addPhoto model.Photo) (model.Photo,error){
	tx:= pr.db.Create(&addPhoto)
	
	return addPhoto, tx.Error
}

func (pr *PhotoRepository) GetAllPhoto () ([]model.Photo,error){
	var photos = []model.Photo{}
	tx:= pr.db.Find(&photos)
	
	return photos, tx.Error	
}
func (pr *PhotoRepository) GetPhotoByUserId (userId string) ([]model.Photo,error){
	var photos = []model.Photo{}
	tx:= pr.db.Find(&photos,"user_id= ?",userId)
	
	return photos, tx.Error	
}

func (pr *PhotoRepository) GetPhotoByPhotoId (id string) (model.Photo,error){
	var photo model.Photo
	tx:= pr.db.First(&photo,"id = ?",id)
	
	return photo, tx.Error	
}


func (pr *PhotoRepository) UpdatePhoto (updatePhoto model.Photo, photoId string) (model.Photo,error){
	tx:= pr.db.Model(&updatePhoto).Where("id = ?",photoId).Updates(&updatePhoto)

	return updatePhoto, tx.Error	
}




// Is it possible to run it ?
func (pr *PhotoRepository) DeletePhoto (id string) (error){
	var deletePhoto model.Photo
	tx:= pr.db.Where("id = ?",id).Delete(&deletePhoto)

	return tx.Error	
}


