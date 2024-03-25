package model

import "time"


type Comment struct{
	ID 		string `gorm:"primaryKey,not null" json:"id"`
	UserID 	string `json:"user_id"`
	PhotoID 	string `json:"photo_id"`
	Message 	string `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
