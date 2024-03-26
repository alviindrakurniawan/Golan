package controller

import (
	"MyGram/dto"
	"MyGram/service"
	"MyGram/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentController struct {
	commentService service.ICommentService
}

func NewCommentController (commentService service.ICommentService) *CommentController{
	return &CommentController{
		commentService: commentService,
	}
}

func (cc *CommentController) AddComment(ctx *gin.Context){
	var addComment dto.AddCommentRequest

	err:= ctx.ShouldBindJSON(&addComment)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data: nil,
		})
		return
	}

	claims,exist:= ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
	}

	sub,err:= util.GetSubFromClaims(claims)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
	}

	comment,err:= cc.commentService.AddComment(addComment,sub.(string),addComment.PhotoID)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated,dto.Response{
		Success: true,
		Message: "Comment added",
		Data: dto.AddCommentResponse{
			ID: comment.ID,
			Message: comment.Message,
			PhotoID: comment.PhotoID,
			UserID: comment.UserID,
			CreatedAt: comment.CreatedAt,
		},
	})

}

func (cc *CommentController) GetCommentById(ctx *gin.Context){
	var commentId = ctx.Param("commentId")
	if commentId==""{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Invalid request",
			Data: nil,
		})
		return
	}
	commentResponse,err:= cc.commentService.GetCommentById(commentId)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK,commentResponse)

}

func (cc *CommentController) UpdateComment (ctx *gin.Context){
	var commmentId = ctx.Param("commentId")
	if commmentId == ""{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Invalid request",
			Data: nil,
		})
		return
	}
	var updateComment dto.UpdateCommentRequest
	err:= ctx.ShouldBindJSON(&updateComment)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data: nil,
		})
		return
	
	}

	//validate request
	errMessage:= validateComment(updateComment)
	if errMessage != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: errMessage,
			Data: nil,
		})
		return
	}
	
	claims,exist:= ctx.Get("claims")
	if !exist{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}
	sub,err:= util.GetSubFromClaims(claims)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}

	comment,err:= cc.commentService.UpdateComment(updateComment,commmentId,sub.(string))
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK,dto.Response{
		Success: true,
		Message: "Comment updated",
		Data: dto.UpdateCommentResponse{
			ID: comment.ID,
			Message: comment.Message,
			PhotoID: comment.PhotoID,
			UserID: comment.UserID,
			UpdateAt: comment.UpdatedAt,
		},
	})
}

func (cc *CommentController) DeleteComment (ctx *gin.Context){
	var commmentId = ctx.Param("commentId")
	if commmentId == ""{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Invalid request",
			Data: nil,
		})
		return
	}

	claims,exist:= ctx.Get("claims")
	if !exist{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}
	sub,err:= util.GetSubFromClaims(claims)
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: "Unauthorized",
			Data: nil,
		})
		return
	}

	err= cc.commentService.DeleteComment(commmentId,sub.(string))
	if err!= nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,dto.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	
	}
	ctx.JSON(http.StatusOK,dto.Response{
		Success: true,
		Message: "Comment deleted",
		Data: dto.DeleteCommentResponse{
			Message: "Comment has been deleted",
		},
	})
}







func validateComment(commentAddOrUpdate interface{}) []string {

	var validate = validator.New()
	err := validate.Struct(commentAddOrUpdate)
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