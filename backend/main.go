package main

import (
 //   "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
//    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

   // "to-do-lister/services"
    "to-do-lister/models"
)

func main() {
    // Connection to PostgreSQL db
    dsn := "user=postgres password=ghq92DAU712.9dn dbname=todolister host=localhost port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // Models automigration
    err = db.AutoMigrate(
        &models.Task{},
        &models.TaskSubTask{},
        &models.Event{},
        &models.EventSubTask{},
        &models.Tag{},
    )
    if err != nil {
        log.Fatal("failed to migrate:", err)
    }

    // Chi router
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

   
    // Run server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Println("Server running on port", port)
    http.ListenAndServe(":"+port, r)
}
