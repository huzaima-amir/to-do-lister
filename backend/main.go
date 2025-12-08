package main

import (
	"database/sql"
	"fmt"
	//"errors"
	//"fmt"
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

// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time  //automatically managed by GORM  for creation time and update time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`  //helps delete stuff without removing from db (softdelete)
}

type User struct {
  ID           uint           // Standard field for the primary key
  Name         string         // A regular string field
  Email        *string        // A pointer to a string, allowing for null values
  Age          uint8          // An unsigned 8-bit integer
  Birthday     *time.Time     // A pointer to time.Time, can be null
  MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
  ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
  CreatedAt    time.Time      // Automatically managed by GORM for creation time
  UpdatedAt    time.Time      // Automatically managed by GORM for update time
  ignored      string         // fields that aren't exported are ignored
}

type APIUser struct {
  ID   uint
  Name string
  TimeAPIUsed *time.Time
}

type Author struct {
  Name  string
  Email string
}

type Blog struct {
  Author
  ID      int
  Upvotes int32
}
// equals
//type Blog struct {
// ID      int64
// Name    string
// Email   string
// Upvotes int32
// }

// equals
// type Blog struct {
//  ID      int
  // Author  Author `gorm:"embedded"`
 // Upvotes int32
// }


// examples/query.go
type Query[T any] interface {
  // SELECT * FROM @@table WHERE id=@id
  GetByID(id int) (T, error)
}

func main() {
    db, err := connectToPostgreSQL()
    if err != nil {
        log.Fatal(err)
    }

    db.Begin()
    db.AutoMigrate(&User{}, &Author{}) // include or data manipulation wont work.

    t := time.Now()

    ////
    user := User{Name: "Jinzhu", Age: 18, Birthday: &t}
  
    result := db.Create(&user)
    if result.Error != nil {
        log.Fatal(result.Error)
    }

    fmt.Println("Inserted user ID:", user.ID)

    db.Select("Name", "Age", "CreatedAt").Create(&user)

    users := []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
    db.Create(&users)

/////
    author := Author(Name: "JKROLLINGINTHEDEEP", Email: "jk@deep.com")
    result2 := db.Create(&author)
    
    if result2.Error != nil {
      log.Fatal(result2.Error)
    }
    
    fmt.Println("Inserted author ID:")




    

    db.Commit()
 }

