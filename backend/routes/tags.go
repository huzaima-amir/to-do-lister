package routes

import (
    "to-do-lister/handlers"
    "to-do-lister/middleware"

    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
)

func TagRoutes(r chi.Router, db *gorm.DB) {

    //  tag routes require authentication
    r.Use(middleware.AuthMiddleware)

    //create (add delete later: TODO!!!)
    r.Post("/", handlers.CreateTagHandler(db))         // create a tag


    //tagging tasks
    r.Post("/tasks/{taskID}", handlers.AddTagToTaskHandler(db))
    r.Delete("/tasks/{taskID}", handlers.RemoveTagFromTaskHandler(db))

    //tagging events
    r.Post("/events/{eventID}", handlers.AddTagToEventHandler(db))
    r.Delete("/events/{eventID}", handlers.RemoveTagFromEventHandler(db))
}
