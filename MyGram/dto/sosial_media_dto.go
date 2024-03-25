package dto

import (
	"time"
)

type NewSosialMediaRequest struct{
	Name string `validate:"required" json:"name"`
	SosialMediaUrl string `validate:"required" json:"sosial_media_url"`
}

type NewSosialMediaResponse struct{
	ID string `json:"id"`
	Name string `json:"name"`
	SosialMediaUrl string `json:"sosial_media_url"`
	UserID string `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

type GetSosialMediaResponse struct{
	SosialMedias []SocialMediaDetail `json:"sosial_medias"`
}

type SocialMediaDetail struct{
	ID string `json:"id"`
	Name string `json:"name"`
	SosialMediaUrl string `json:"sosial_media_url"`
	UserID string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User userSosialMedia `json:"user"`

}

type userSosialMedia struct{
	ID string `json:"id"`
	UserName string `json:"user_name"`
	ProfileImageUrl string `json:"profile_image_url"`
}



type UpdateSosialMediaRequest struct{
	Name string `json:"name"`
	SosialMediaUrl string `json:"sosial_media_url"`
}

type UpdateSosialMediaResponse struct{
	ID string `json:"id"`
	Name string `json:"name"`
	SosialMediaUrl string `json:"sosial_media_url"`
	UserID string `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteSosialMediaResponse struct{
	Message string `json:"message"`
}