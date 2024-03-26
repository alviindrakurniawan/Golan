package dto

import "time"


type AddCommentRequest struct{
	Message 	string `validate:"required" json:"message"`
	PhotoID 	string `validate:"required" json:"photo_id"`
}

type AddCommentResponse struct{
	ID 		string `json:"id"`
	Message 	string `json:"message"`
	PhotoID 	string `json:"photo_id"`
	UserID 	string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}


type GetCommentResponse struct{
	ID 		string `json:"id"`
	Message 	string `json:"message"`
	PhotoID 	string `json:"photo_id"`
	UserID 	string `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`	
	UserComment UserComment `json:"user"`
	PhotoComment PhotoComment `json:"photo"`
}

type UserComment struct{
	ID string `json:"id"`
	Email string `json:"email"`
	UserName string `json:"user_name"`
}

type PhotoComment struct{
	ID        string    `json:"id" `
	Title     string    `json:"title" `
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
}

type UpdateCommentRequest struct{
	Message 	string `validate:"required" json:"message"`
}

type UpdateCommentResponse struct{
	ID 		string `json:"id"`
	Message 	string `json:"message"`
	PhotoID 	string `json:"photo_id"`
	UserID 	string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct{
	Message string `json:"message"`
}