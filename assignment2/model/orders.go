package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Orders struct {
	ID string `json:"id" gorm:"primaryKey"`
	CustomerName string `json:"customerName" gorm:"not null"`
	Items []Item `gorm:"foreignKey:OrdersID"`
	OrderedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time 
	DeletedAt gorm.DeletedAt 
	
}

func (o *Orders) BeforeCreate(tx *gorm.DB)error{
	o.ID = uuid.New().String()
	o.OrderedAt = time.Now()

	return nil
}

func (o *Orders) BeforeUpdate(tx *gorm.DB)error{
	o.UpdatedAt = time.Now()

	return nil
}