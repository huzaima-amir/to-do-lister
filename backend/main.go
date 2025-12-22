package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "to-do-lister/handlers"
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

    // Routes
    r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
        var input struct {
            Title       string    `json:"title"`
            Description string    `json:"description"`
            Deadline    time.Time `json:"deadline"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        id := handlers.CreateTask(db, input.Title, input.Description, input.Deadline)

        json.NewEncoder(w).Encode(map[string]interface{}{
            "id": id,
        })
    })

    r.Put("/tasks/{id}/start", func(w http.ResponseWriter, r *http.Request) {
        idStr := chi.URLParam(r, "id")
        var id uint
        fmt.Sscan(idStr, &id)

        if err := handlers.StartTask(db, id); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{
            "status": "started",
        })
    })

    r.Put("/tasks/{id}/end", func(w http.ResponseWriter, r *http.Request) {
        idStr := chi.URLParam(r, "id")
        var id uint
        fmt.Sscan(idStr, &id)

        if err := handlers.EndTask(db, id); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{
            "status": "finished",
        })
    })

    r.Delete("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
        idStr := chi.URLParam(r, "id")
        var id uint
        fmt.Sscan(idStr, &id)

        handlers.DeleteTask(db, id)

        json.NewEncoder(w).Encode(map[string]string{
            "status": "deleted",
        })
    })

    // Run server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Println("Server running on port", port)
    http.ListenAndServe(":"+port, r)
}
