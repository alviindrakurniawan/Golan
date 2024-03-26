package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Comment struct{
	ID 		string `gorm:"primaryKey,not null" json:"id"`
	UserID 	string `json:"user_id"`
	PhotoID 	string `json:"photo_id"`
	Message 	string `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (c *Comment) BeforeCreate(tx *gorm.DB)error{
	c.ID= uuid.New().String()
	c.CreatedAt = time.Now()

	return nil
}