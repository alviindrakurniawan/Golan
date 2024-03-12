package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Item struct {
	ItemId string `gorm:"primaryKey"`
	ItemCode string `gorm:"type:varchar(100) not null"`
	Description string `gorm:"type:varchar(200) not null"`
	Quantity int
	OrdersID string

}

func (i *Item) BeforeCreate(tx *gorm.DB)error{
	i.ItemId = uuid.New().String()

	return nil
}
