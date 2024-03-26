package repository

import (
	"MyGram/model"

	"gorm.io/gorm"
)


type ICommentRepository interface {
	AddComment(addComment model.Comment) (model.Comment,error)
	GetComment(id string) (model.Comment,error)
	GetCommentByUserId(userId string) ([]model.Comment,error)
	UpdateComment(updateComment model.Comment ,commendId string) (model.Comment,error)
	DeleteComment(commentId string) (error)
}

type CommentRepository struct{
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB)(*CommentRepository){
	return &CommentRepository{db}

}

func (cr *CommentRepository) AddComment(addComment model.Comment) (model.Comment,error){
	tx:= cr.db.Create(&addComment)
	
	return addComment, tx.Error	
}

func (cr *CommentRepository) GetComment(id string) (model.Comment,error){
	var comment model.Comment
	tx:= cr.db.Find(&comment,"id = ?",id)

	return comment,tx.Error
}

func (cr *CommentRepository) GetCommentByUserId(userId string) ([]model.Comment,error){
	var comment []model.Comment
	tx:= cr.db.Find(&comment,"user_id = ?",userId)

	return comment,tx.Error
}

func (cr *CommentRepository) UpdateComment(updateComment model.Comment ,commendId string) (model.Comment,error){
	tx:= cr.db.Model(&updateComment).Where("id = ?",commendId).Updates(&updateComment)
	
	return updateComment,tx.Error
}

func (cr *CommentRepository) DeleteComment(commentId string) (error){
	var deleteComment model.Comment
	tx:= cr.db.Where("id =?",commentId).Delete(&deleteComment)
	
	return tx.Error
}