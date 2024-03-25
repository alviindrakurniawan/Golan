package repository

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type SocialMediaRepository struct{
	db * gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB)(*SocialMediaRepository){
	return &SocialMediaRepository{db}
}


func (sr *SocialMediaRepository) AddSosialMedia (addSosialMedia model.SosialMedia) (model.SosialMedia,error){
	tx:= sr.db.Create(&addSosialMedia)
	
	return addSosialMedia, tx.Error	
}

func (sr *SocialMediaRepository) GetSosialMedia (id string) ([]model.SosialMedia,error){
	var sosialMedia []model.SosialMedia
	tx:= sr.db.Find(&sosialMedia,"id = ?",id)

	return sosialMedia, tx.Error
}

func (sr *SocialMediaRepository) UpdateSosialMedia (updateSosialMedia model.SosialMedia,id string) (model.SosialMedia,error){
	tx:= sr.db.Model(&updateSosialMedia).Where("id = ?",id).Updates(&updateSosialMedia)

	return updateSosialMedia, tx.Error
}


