package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
//	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"to-do-lister/models"
	"to-do-lister/routes"
    
    chimw "github.com/go-chi/chi/v5/middleware"  // golang chi package middleware
    custommw  "to-do-lister/middleware"        //custom middleware
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

    //middleware
    r.Use(custommw.Cors())
    r.Use(chimw.Logger)
    r.Use(chimw.Recoverer)

    // Mount route groups
    r.Route("/tasks", func(tr chi.Router) { routes.TaskRoutes(tr, db) }) 
    r.Route("/events", func(er chi.Router) { routes.EventRoutes(er, db) }) 
    r.Route("/tags", func(tg chi.Router) { routes.TagRoutes(tg, db) }) 
    r.Route("/users", func(ur chi.Router) { routes.UserRoutes(ur, db) })

    //start server
    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", r)

}
