package controller

import (
	"MyGram/dto"
	"MyGram/service"
	"MyGram/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var newUser dto.RegisterRequest

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	fmt.Println(newUser)
	errMessage := validateRegister(newUser)
	if errMessage != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: errMessage,
			Data:    nil,
		})
		return
	}

	hashedPassword, err := util.HashPassword([]byte(newUser.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	newUser.Password = string(hashedPassword)

	registeredUser, err := uc.UserService.RegisterUser(newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return

	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "User registered successfully",
		Data: dto.RegisterResponse{
			Age:      registeredUser.Age,
			Email:    registeredUser.Email,
			ID:       registeredUser.ID,
			UserName: registeredUser.UserName,
		},
	})
}

func (us *UserController) Login(ctx *gin.Context) {
	var loginUser dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}


	token,err:= us.UserService.Login(loginUser)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK,dto.Response{
		Success: true,
		Message: "User logged in successfully",
		Data: token,
	})
}

func (us *UserController) UpdateUser(ctx *gin.Context) {
	var param = ctx.Param("id")
	if param == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request",
			Data:    nil,
		})
		return
	}

	claims, exist := ctx.Get("claims")
	if !exist{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}

	sub,err:= util.GetSubFromClaims(claims)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}

	if sub != param{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	
	}

	var updateUser dto.UpdateUserRequest
	err = ctx.ShouldBindJSON(&updateUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	user,err:= us.UserService.UpdateUser(updateUser)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}
	
	ctx.JSON(http.StatusOK,dto.Response{
		Success: true,
		Message: "User updated successfully",
		Data: dto.UpdateUserResponse{
			ID: user.ID,
			UserName: user.UserName,
			Email: user.Email,
			Age: user.Age,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

func (uc *UserController) DeleteUser(ctx *gin.Context){
	claims, exist := ctx.Get("claims")
	if !exist{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
	}

	sub,err:= util.GetSubFromClaims(claims)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
	}

	err= uc.UserService.DeleteUser(sub.(string))
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
	}

	ctx.JSON(http.StatusOK,dto.Response{
		Message: "Your account has been successfully deleted",
	})

}



func validateRegister(newUser dto.RegisterRequest) []string {
	var validate = validator.New()
	err := validate.Struct(newUser)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		log.Println(errs)
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, e.Field()+" is required")
			case "min":
				errorMessage = append(errorMessage, fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param()))
			case "email":
				errorMessage = append(errorMessage, fmt.Sprintf("%s format is invalid", e.Field()))
			case "unique":
				errorMessage = append(errorMessage, fmt.Sprintf("%s already exists", e.Field()))
			}
		}
		return errorMessage
	}
	return nil
}

	
	
	
