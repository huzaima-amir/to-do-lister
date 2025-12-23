package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"to-do-lister/models"
	"to-do-lister/routes"
)

func main() {

    // Connection to postgres db
    dsn := "user=postgres password=ghq92DAU712.9dn dbname=todolister host=localhost port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // models automigration
    err = db.AutoMigrate( 
        &models.User{},
        &models.Task{},
        &models.TaskSubTask{},
        &models.Event{},
        &models.EventSubTask{},
        &models.Tag{},
    )
    if err != nil {
        log.Fatal("failed to migrate:", err)
    }

    // Create router
    r := chi.NewRouter()

    // Global middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)


    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    }))

    // Mount route groups
    r.Route("/users", func(ur chi.Router) {
        routes.UserRoutes(ur, db)
    })

    //start server
    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", r)

}
