package controller

import (
	"assignment2/model"
	"assignment2/repository"
	"assignment2/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *OrderController{
	
	return &OrderController{
		orderRepository: orderRepository,
	}
}

func (oc *OrderController)CreateOrder (ctx * gin.Context){
	var newOrder model.Orders

	err:= ctx.ShouldBindJSON(&newOrder)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,util.PrintResponse(false,nil,err.Error()))
		return
	}

	createdOrder,err:=oc.orderRepository.CreateOrder(newOrder)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,util.PrintResponse(false,nil,err.Error()))
		return
	}

	ctx.JSON(http.StatusOK,util.PrintResponse(true,createdOrder,""))

}

func (oc *OrderController)GetOrder (ctx *gin.Context){
	var orders []model.Orders

	orders,err:= oc.orderRepository.GetOrder(orders)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,util.PrintResponse(false,nil,err.Error()))
		return
	}

	ctx.JSON(http.StatusOK,util.PrintResponse(true,orders,""))

}

func (oc *OrderController)UpdateOrder (ctx *gin.Context){
	var updatedOrder model.Orders
	idString:= ctx.Param("id")


	err:= ctx.ShouldBindJSON(&updatedOrder)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,util.PrintResponse(false,nil,err.Error()))
		return
	}

	updatedOrder,err= oc.orderRepository.UpdateOrder(updatedOrder,idString)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,util.PrintResponse(false,nil,err.Error()))
		return
	
	}

	ctx.JSON(http.StatusOK,util.PrintResponse(true,updatedOrder,""))
}

func (oc *OrderController)DeleteOrder (ctx *gin.Context){
	// targetOrder := model.Orders{}
	var idString= ctx.Param("id")

	targetOrder,err:= oc.orderRepository.DeleteOrder(idString)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,util.PrintResponse(false,nil,err.Error()))
		return
	}

	message := fmt.Sprintf("Deleted Account %s With Name %s Success", targetOrder.ID, targetOrder.CustomerName)

	ctx.JSON(http.StatusOK,util.PrintResponse(true,message,""))
}

