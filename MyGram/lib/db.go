package lib

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDatabase() (*gorm.DB,error ){
	host:= "localhost"
	port:= 5432
	dbUser:= "postgres"
	dbPassword:= "password"
	dbName:= "mygram"


	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, dbUser, dbPassword, dbName)
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}