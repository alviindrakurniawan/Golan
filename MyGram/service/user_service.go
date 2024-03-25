package service

import (
	"MyGram/dto"
	"MyGram/model"
	"MyGram/repository"
	"MyGram/util"
	"errors"
)
type IUserService interface{
	RegisterUser(newUser dto.RegisterRequest)(model.User,error)
	Login(loginUser dto.LoginRequest)(dto.LoginResponse,error)
	UpdateUser(updateUser dto.UpdateUserRequest)(model.User,error)
	DeleteUser(userId string) error
}


type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService (userRepository repository.IUserRepository) *UserService {
	
	return &UserService{
		userRepository:userRepository,
	}
}

func (us *UserService) RegisterUser (newUser dto.RegisterRequest)(model.User,error){


	var user=  model.User{
		UserName: newUser.UserName,
		Email: newUser.Email,
		Password: newUser.Password,
		Age: newUser.Age,
	}
	
	
	registeredUser,err:= us.userRepository.RegisterUser(user)
	if err != nil{
		return model.User{},err
	}

	return registeredUser,nil
}

func (us *UserService) Login (loginUser dto.LoginRequest) (dto.LoginResponse,error){
	
	user,err:= us.userRepository.GetUserByEmail(loginUser.Email)
	if err != nil{
		
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	ok := util.HashMacthesPassword([]byte(user.Password),[]byte(loginUser.Password))
	if !ok{
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	token,err:= util.GenerateJWTToken(user.ID)
	if err != nil{
		return dto.LoginResponse{},err
	}

	return dto.LoginResponse{Token: token},nil

}

func (us *UserService) UpdateUser(updateUser dto.UpdateUserRequest)(model.User,error){
	
	user,err := us.userRepository.GetUserByEmail(updateUser.Email)
	if err != nil{
		return model.User{},errors.New("user not found")
	}

	user.UserName = updateUser.UserName

	updatedUser,err:= us.userRepository.UpdateUser(user,user.ID)
	if err != nil{
		return model.User{},err
	}

	return updatedUser,nil

}

func (us *UserService) DeleteUser(userId string) error{
	user,err:= us.userRepository.GetUserById(userId)
	if err != nil{
		return errors.New("user not found")
	}

	err= us.userRepository.DeleteUser(user)
	if err != nil{
		return err
	}

	return nil
}



