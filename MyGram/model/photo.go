package model

import "time"

type Photo struct{
	ID        string    `gorm:"primaryKey,not null" json:"id" `
	Title     string    `json:"title" `
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
