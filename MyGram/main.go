package main

import (
	"MyGram/controller"
	"MyGram/lib"
	"MyGram/model"
	"MyGram/repository"
	"MyGram/service"

	"github.com/gin-gonic/gin"
)

func main (){
	db,err:=lib.InitDatabase()
	if err != nil{
		panic(err)
	}
	
	err = db.AutoMigrate(&model.User{},&model.SosialMedia{},&model.Comment{},&model.Photo{})
	if err != nil{
		panic(err)
	}

	userRepository:= repository.NewUserRepository(db)
	userService:= service.NewUserService(userRepository)
	userController:= controller.NewUserController(userService)
	
	ginEngine:= gin.Default()
	ginEngine.POST("/users/register",userController.RegisterUser)
	ginEngine.POST("/users/login",userController.Login)
	ginEngine.PUT("/users",userController.UpdateUser)
	ginEngine.DELETE("/users",userController.DeleteUser)


	err= ginEngine.Run("localhost:8080")
	if err!= nil{
		panic(err)
	}
}