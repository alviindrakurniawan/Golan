package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct{
	ID        string    `gorm:"primaryKey,not null" json:"id" `
	Title     string    `json:"title" `
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	Comments  []Comment `json:"comments"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) BeforeCreate (tx *gorm.DB)error{
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()

	return nil
}
