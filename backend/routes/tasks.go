package routes

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "gorm.io/gorm"

    "to-do-lister/services"
)

func NewRouter(db *gorm.DB) *chi.Mux {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

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

        id := services.CreateTask(db, input.Title, input.Description, input.Deadline,9)
        json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
    })

    r.Put("/tasks/{id}/start", func(w http.ResponseWriter, r *http.Request) {
        idStr := chi.URLParam(r, "id")
        var id uint
        fmt.Sscan(idStr, &id)

        if err := services.StartTask(db, id); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{"status": "started"})
    })

    r.Put("/tasks/{id}/end", func(w http.ResponseWriter, r *http.Request) {
        idStr := chi.URLParam(r, "id")
        var id uint
        fmt.Sscan(idStr, &id)

        if err := services.EndTask(db, id); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{"status": "finished"})
    })

    r.Delete("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
        idStr := chi.URLParam(r, "id")
        var id uint
        fmt.Sscan(idStr, &id)

        services.DeleteTask(db, id)
        json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
    })

    return r
}
