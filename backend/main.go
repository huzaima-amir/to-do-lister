package main

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)

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
 
}
