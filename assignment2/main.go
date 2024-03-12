package main

import (
	"assignment2/controller"
	"assignment2/lib"
	"assignment2/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main (){

	db,err:=lib.StartDB()
	if err!=nil{
		log.Printf("Error database : %v",err)
	}

	orderRepository:= repository.NewOrderRepository(db)
	orderController:= controller.NewOrderController(orderRepository)

	ginEngine:= gin.Default()

	ginEngine.POST("/order",orderController.CreateOrder)
	ginEngine.GET("/order",orderController.GetOrder)
	ginEngine.PUT("/order/:id",orderController.UpdateOrder)
	ginEngine.DELETE("/order/:id",orderController.DeleteOrder)
	


	err = ginEngine.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

}