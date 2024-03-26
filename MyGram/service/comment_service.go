package service

import (
	"MyGram/dto"
	"MyGram/model"
	"MyGram/repository"
	"errors"
)

type ICommentService interface {
	AddComment(addComment dto.AddCommentRequest,userId string,photoId string)(model.Comment,error)
	GetCommentById(id string)(dto.GetCommentResponse,error)
	UpdateComment(updateComment dto.UpdateCommentRequest, commentId string,userId string)(model.Comment,error)
	DeleteComment(commentId string,userId string)error
}

type CommentService struct {
	commentRepository repository.ICommentRepository
	photoRepository repository.IPhotoRepository 
	userRepositoy repository.IUserRepository

}

func NewCommentService (commentRepository repository.ICommentRepository, photoRepository repository.IPhotoRepository, userRepositoy repository.IUserRepository ) *CommentService{
	return &CommentService{
		commentRepository: commentRepository,
		photoRepository: photoRepository,
		userRepositoy: userRepositoy,
	}
}

func (cs *CommentService) AddComment (addComment dto.AddCommentRequest,userId string,photoId string)(model.Comment,error){

	var comment= model.Comment{
		Message: addComment.Message,
		UserID: userId,
		PhotoID: photoId,
	}
	addedComment,err:= cs.commentRepository.AddComment(comment)
	if err != nil{
		return model.Comment{},err
	}
	return addedComment,nil

}

// Get 1 comment
func (cs *CommentService) GetCommentById(id string)(dto.GetCommentResponse,error){
	comment,err:= cs.commentRepository.GetComment(id)
	if err != nil{
		return dto.GetCommentResponse{},err
	}

	userComment,err:= cs.userRepositoy.GetUserById(comment.UserID)
	if err != nil{
		return dto.GetCommentResponse{},err
	}

	photoComment,err:= cs.photoRepository.GetPhotoByPhotoId(comment.PhotoID)
	if err!= nil{
		return dto.GetCommentResponse{},err
	}

	return dto.GetCommentResponse{
		ID: comment.ID,
		Message: comment.Message,
		PhotoID: comment.PhotoID,
		UserID: comment.UserID,
		UpdatedAt: comment.UpdatedAt,
		CreatedAt: comment.CreatedAt,
		UserComment: dto.UserComment{
			ID: userComment.ID,
			Email: userComment.Email,
			UserName: userComment.UserName,
		},
		PhotoComment: dto.PhotoComment{
			ID: photoComment.ID,
			Title: photoComment.Title,
			Caption: photoComment.Caption,
			PhotoUrl: photoComment.PhotoUrl,
			UserID: photoComment.UserID,
		},
	},nil
		
	
}

func (cs *CommentService) UpdateComment (updateComment dto.UpdateCommentRequest, commentId string,userId string)(model.Comment,error){
	currentComments,err:= cs.commentRepository.GetCommentByUserId(userId)
	if err != nil{
		return model.Comment{},err
	}

	var commentVerified bool = false
	for _,comment:= range currentComments{
		if comment.ID == commentId{
			commentVerified = true
		}
	}
	
	if !commentVerified{
		return model.Comment{},errors.New("Unauthorized")
	}
	

	var updatecomment = model.Comment{
		Message: updateComment.Message,
	}

	updatedComment,err:= cs.commentRepository.UpdateComment(updatecomment,commentId)
	if err != nil{
		return model.Comment{},err
	}
	
	return updatedComment,nil
}

func (cs *CommentService) DeleteComment(commentId string,userId string)error{
	deleteComment,err:= cs.commentRepository.GetComment(commentId)
	if err != nil{
		return errors.New("comment not found")
	}

	if deleteComment.UserID != userId {
		return errors.New("Unauthorized")
	}

	err= cs.commentRepository.DeleteComment(deleteComment.ID)
	if err != nil{
		return err
	}

	return nil
}





