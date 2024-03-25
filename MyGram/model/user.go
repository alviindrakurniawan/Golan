package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct{
	ID string `gorm:"primaryKey,unique,not null" json:"id"`
	UserName string `gorm:"uniqueIndex" json:"user_name"`
	Email string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password"`
	Age int `json:"age"`
	Photos []Photo `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt 
}


func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error{
	u.UpdatedAt= time.Now()
	return nil

}