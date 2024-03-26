package repository

import (
	"MyGram/model"

	"gorm.io/gorm"
)



type IUserRepository interface {
	RegisterUser(newUser model.User) (model.User,error)
	GetUserById(id string) (model.User,error)
	GetUserByEmail(email string) (model.User,error)
	UpdateUser(updateUser model.User,id string) (model.User,error)
	DeleteUser(model.User) (error)

}

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) RegisterUser(newUser model.User) (model.User,error){
	tx:= ur.db.Create(&newUser)
	
	return newUser, tx.Error	
}

func (ur *UserRepository) GetUserById(id string) (model.User,error){
	var user model.User
	tx:= ur.db.First(&user,"id = ?", id)
	
	return user, tx.Error	
}

func (ur *UserRepository) GetUserByEmail(emailUser string) (model.User,error){
	var user model.User
	tx:= ur.db.First(&user,"email= ?", emailUser)
	
	return user, tx.Error	
}

func (ur *UserRepository) UpdateUser(updateUser model.User,id string) (model.User,error){
	tx:= ur.db.Model(&updateUser).Where("id = ?",id).Updates(&updateUser)
	
	
	return updateUser, tx.Error	
}


func (ur *UserRepository) DeleteUser(deleteUser model.User) (error){
	userId:=deleteUser.ID
	tx:= ur.db.Where("id =?",userId).Delete(&deleteUser)
	
	return tx.Error	
}

