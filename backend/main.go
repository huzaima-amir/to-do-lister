package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connection
func connectToPostgreSQL() (*gorm.DB, error) {  
    dsn := "user=postgres password=ghq92DAU712.9dn dbname=todolister host=localhost port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}


//  entities:
type Task struct {
  //starting time
  title string
  description string
  // deadline time
  started bool
  finished bool
  checklist [] string
 }

type Event struct {
  title string
  description string
  startingTime time.Time
  endTime time.Time
  place string
  checklist[] string
}

type Tag struct { // to describe events and tasks?
  title string
  description string 
}

type checklist struct {
  subtasks []string
  totalsubtasks, finishedsubtasks int
}

func main() {
    db, err := connectToPostgreSQL()
    if err != nil {
        log.Fatal(err)

    }

    db.Begin()
 
}
