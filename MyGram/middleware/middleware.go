package middleware

import (
	"MyGram/dto"
	"MyGram/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context){
	authorization:= ctx.GetHeader("Authorization")
	splittedAuthorization:= strings.Split(authorization,"Bearer ")
	if (len(splittedAuthorization)<=1){
		var response = dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,response)
		return
	}

	token:= splittedAuthorization[1]
	claims,err:= util.GetJWTClaims(token)
	if err!=nil{
		var response = dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,response)
		return
	}

	ctx.Set("claims",claims)
	ctx.Next()

}