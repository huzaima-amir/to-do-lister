package main

import (
	//	"fmt"
	//"errors"
	//"fmt"
///	"database/handlers"
	"database/models"
	//"fmt"
	"log"
	//"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connection to Database
func connectToPostgreSQL() (*gorm.DB, error) {  
    dsn := "user=postgres password=ghq92DAU712.9dn dbname=todolister host=localhost port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

func main() {
    db, err := connectToPostgreSQL()
    if err != nil {
        log.Fatal(err)
    }
	db.Begin()

   db.AutoMigrate(&models.Task{},&models.Tag{})
 }

