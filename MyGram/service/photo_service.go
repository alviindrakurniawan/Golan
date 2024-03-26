package service

import (
	"MyGram/dto"
	"MyGram/model"
	"MyGram/repository"
	"errors"
)

type IPhotoService interface {
	AddPhoto(addPhoto dto.AddPhotoRequest,userId string)(model.Photo,error)
	GetPhoto()([]dto.GetPhotoResponse,error)
	UpdatePhoto(updatePhoto dto.UpdatePhotoRequest,userId string, photoId string)(model.Photo,error)
	DeletePhoto(photoId string,userId string) error
}
type PhotoService struct {
	photoRepository repository.IPhotoRepository
	userRepository repository.IUserRepository
}

func NewPhotoService (photoRepository repository.IPhotoRepository, userRepository repository.IUserRepository) *PhotoService{
	return &PhotoService{
		photoRepository: photoRepository,
		userRepository: userRepository,
	}
}

func (ps *PhotoService) AddPhoto (addPhoto dto.AddPhotoRequest,userId string)(model.Photo,error){


	var photo= model.Photo{
		PhotoUrl: addPhoto.PhotoUrl,
		UserID: userId,
		Caption: addPhoto.Caption,
		Title: addPhoto.Title,
	}
	addedPhoto,err:= ps.photoRepository.AddPhoto(photo)
	if err != nil{
		return model.Photo{},err
	}
	return addedPhoto,nil
}

func (ps *PhotoService) GetPhoto ()([]dto.GetPhotoResponse,error){
	photos,err:= ps.photoRepository.GetAllPhoto()
	if err != nil{
		return []dto.GetPhotoResponse{},err
	}
	var allPhoto []dto.GetPhotoResponse
	for _,photo:= range photos{
		user,err := ps.userRepository.GetUserById(photo.UserID)
		if err != nil{
			return []dto.GetPhotoResponse{},err
		}

		allPhoto = append(allPhoto,dto.GetPhotoResponse{
			ID: photo.ID,
			Title: photo.Title,
			Caption: photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserID: photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: dto.UserPhoto{
				Email: user.Email,
				UserName: user.UserName,
			},
		})
	}

	return allPhoto,nil

}

func (ps *PhotoService) UpdatePhoto(updatePhoto dto.UpdatePhotoRequest,userId string, photoId string)(model.Photo,error){
	
	photos_user,err:= ps.photoRepository.GetPhotoByUserId(userId)
	if err != nil{
		return model.Photo{},err
	}

	//verify photo
	var photoVerified bool = false
	for _,photo:=range photos_user{
		if photo.ID == photoId{
			photoVerified = true
		}
	}	
	if !photoVerified{
		return model.Photo{},errors.New("Unauthorized")
	}

	//update photo
	var photo = model.Photo{
		Title: updatePhoto.Title,
		Caption: updatePhoto.Caption,
		PhotoUrl: updatePhoto.PhotoUrl,
	}
	updatedPhoto,err:=ps.photoRepository.UpdatePhoto(photo,photoId)
	if err!= nil{
		return model.Photo{},err
	}
	

	return updatedPhoto,nil
}

func (ps *PhotoService) DeletePhoto(photoId string,userId string) error {
	photos_user, err := ps.photoRepository.GetPhotoByUserId(userId)
	if err != nil {
		return err
	}	

	//verify photo
	var photoVerified bool = false
	for _, photo := range photos_user {
		if photo.ID == photoId {
			photoVerified = true
		}
	}
	if !photoVerified {
		return errors.New("Unauthorized")
	}

	// //check photo id exist
	// _, err = ps.photoRepository.GetPhotoByPhotoId(photoId)
	// if err != nil {
	// 	return errors.New("photo not found")
	// }
	


	err = ps.photoRepository.DeletePhoto(photoId)
	if err != nil {
		return err
	}

	return nil
}











