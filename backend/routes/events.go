package routes

import (
    "to-do-lister/handlers"
    "to-do-lister/middleware"
    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
)

func EventRoutes(r chi.Router, db *gorm.DB) {

    // event routes require authentication
    r.Use(middleware.AuthMiddleware)

    // CRUD
    r.Get("/", handlers.GetEventsHandler(db))            // list all events for user
    r.Get("/{eventID}", handlers.GetEventByIDHandler(db)) // get single event
    r.Post("/", handlers.CreateEventHandler(db))         // create event
    r.Delete("/{eventID}", handlers.DeleteEventHandler(db)) // delete event

    // Subtasks
    r.Post("/{eventID}/subtasks", handlers.AddSubtasktoEventHandler(db))
    r.Put("/{eventID}/subtasks/{subtaskID}", handlers.ToggleEventSubtaskHandler(db))
    r.Delete("/{eventID}/subtasks/{subtaskID}", handlers.DeleteSubTaskByEventHandler(db))
}

