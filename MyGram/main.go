package main

import (
	"MyGram/controller"
	"MyGram/lib"
	"MyGram/middleware"
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

	photoRepository:= repository.NewPhotoRepository(db)
	photoService:= service.NewPhotoService(photoRepository,userRepository)
	photoControl:= controller.NewPhotoController(photoService)
	
	ginEngine:= gin.Default()

	ginEngine.POST("/users/register",userController.RegisterUser)
	ginEngine.POST("/users/login",userController.Login)


	userGroup := ginEngine.Group("/users", middleware.AuthMiddleware)
	userGroup.PUT("/:Id",userController.UpdateUser)
	userGroup.DELETE("/:Id",userController.DeleteUser)

	photoGroup:= ginEngine.Group("/photos",middleware.AuthMiddleware)
	photoGroup.GET("/",photoControl.GetAllPhoto)
	photoGroup.POST("/",photoControl.AddPhoto)
	photoGroup.PUT("/:photoId",photoControl.UpdatePhoto)
	photoGroup.DELETE("/:photoId",photoControl.DeletePhoto)

	err= ginEngine.Run("localhost:8080")
	if err!= nil{
		panic(err)
	}
}