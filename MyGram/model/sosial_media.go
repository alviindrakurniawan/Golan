package model

type SosialMedia struct{
	ID string `gorm:"primaryKey,not null" json:"id"`
	Name string `json:"name"`
	SosialMediaUrl string `json:"sosial_media_url"`
	UserID string `json:"user_id"`
}