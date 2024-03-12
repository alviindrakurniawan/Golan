package lib

import (
	"assignment2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func StartDB() (*gorm.DB, error) {
	connectionString:= "host=localhost user=postgres password=password dbname=orders_by port=5432 sslmode=disable"

	db,err:= gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err!=nil{
		return nil, fmt.Errorf("failed to connect db: %v", err)
	}

	if err= db.Debug().AutoMigrate(&model.Orders{},&model.Item{});err!=nil{
		return nil, fmt.Errorf("failed to migrate db: %v", err)
	}

	return db,nil
}
