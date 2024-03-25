package dto

import (
	"time"
)

type AddPhotoRequest struct{
	Title     string    `validate:"required" json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `validate:"required" json:"photo_url"`
}

type AddPhotoResponse struct{
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}


type UserPhoto struct{
	Email string `json:"email"`
	UserName string `json:"user_name"`
}

type GetPhotoResponse struct{
	ID        string    `json:"id" `
	Title     string    `json:"title" `
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	User      UserPhoto `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type UpdatePhotoRequest struct{
	Title     string    `json:"title" `
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
}


type UpdatePhotoResponse struct{
	ID        string    `json:"id" `
	Title     string    `json:"title" `
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}


type DeletePhotoResponse struct{
	Message string `json:"message"`
}
