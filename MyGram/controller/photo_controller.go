package controller

import (
	"MyGram/dto"
	"MyGram/service"
	"MyGram/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type PhotoController struct{
	photoService service.IPhotoService
}

func NewPhotoController(photoService service.IPhotoService) *PhotoController{
	return &PhotoController{
		photoService: photoService,
	}
}

func (pc *PhotoController) AddPhoto (ctx *gin.Context){
	var addPhoto dto.AddPhotoRequest

	err:= ctx.ShouldBindJSON(&addPhoto)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data: nil,
		})
		return
 	}

	errMessage:= validatePhoto(addPhoto)
	if errMessage != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: errMessage,
			Data: nil,
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
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}
	// var addedPhoto model.Photo
	addedPhoto,err:= pc.photoService.AddPhoto(addPhoto,sub.(string))
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK,dto.Response{
		Success: true,
		Message: "Photo added successfully",
		Data: dto.AddPhotoResponse{
			ID: addedPhoto.ID,
			Title: addedPhoto.Title,
			Caption: addedPhoto.Caption,
			PhotoUrl: addedPhoto.PhotoUrl,
			UserID: addedPhoto.UserID,
			CreatedAt: addedPhoto.CreatedAt,
		},
	})
}

func (pc *PhotoController) GetAllPhoto(ctx *gin.Context){
	var allphoto []dto.GetPhotoResponse

	allphoto,err:= pc.photoService.GetPhoto()
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK,dto.Response{
		Success: true,
		Message: "Get All Photo Success",
		Data: allphoto,
	})
	
}

func(pc *PhotoController) UpdatePhoto (ctx *gin.Context){
	var photoId = ctx.Param("photoId")
	if photoId == ""{
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
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}

	var updatePhoto dto.UpdatePhotoRequest
	err= ctx.ShouldBindJSON(&updatePhoto)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data: nil,
		})
		return
	}

	errMessage:= validatePhoto(updatePhoto)
	if errMessage != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: errMessage,
			Data: nil,
		})
		return
	}

	updatedPhoto,err:= pc.photoService.UpdatePhoto(updatePhoto,sub.(string),photoId)
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK,dto.Response{
		Success: true,
		Message: "Photo updated successfully",
		Data: dto.UpdatePhotoResponse{
			ID: updatedPhoto.ID,
			Title: updatedPhoto.Title,
			Caption: updatedPhoto.Caption,
			PhotoUrl: updatedPhoto.PhotoUrl,
			UserID: updatedPhoto.UserID,
			UpdatedAt: updatedPhoto.UpdatedAt,
		},
	})

}

func (pc *PhotoController) DeletePhoto(ctx *gin.Context){
	var photoId = ctx.Param("photoId")
	if photoId == ""{
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
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}

	err= pc.photoService.DeletePhoto(photoId,sub.(string))
	if err!=nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}
	
	ctx.JSON(http.StatusOK,dto.Response{
		Message: "Your photo has been successfully deleted",
	})
}


//How to hanlde validation input 2 ?
func validatePhoto(photoAddOrUpdate interface{}) []string {

	var validate = validator.New()
	err := validate.Struct(photoAddOrUpdate)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, e.Field()+" is required")
			}
		}
		return errorMessage
	}
	return nil
}

