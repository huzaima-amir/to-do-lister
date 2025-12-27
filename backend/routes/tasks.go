package routes

import (
	"to-do-lister/handlers"
	"github.com/go-chi/chi/v5"
    "gorm.io/gorm"
    "to-do-lister/middleware"
	)



func TaskRoutes(r chi.Router, db *gorm.DB) {
    r.Use(middleware.AuthMiddleware)

    //CRUD
    r.Get("/", handlers.GetTasksHandler(db))
	r.Get("/{taskID}", handlers.GetTaskByIDHandler(db))
    r.Post("/", handlers.CreateTaskHandler(db))
    r.Delete("/{taskID}", handlers.DeleteTaskHandler(db))

    //state transitions
    r.Put("/{taskID}/start", handlers.StartTaskHandler(db))
    r.Put("/{taskID}/end", handlers.EndTaskHandler(db))

    //Subtasks
    r.Post("/{taskID}/subtasks", handlers.AddSubtasktoTaskHandler(db))
    r.Put("/{taskID}/subtasks/{subtaskID}", handlers.ToggleTaskSubtaskHandler(db))
    r.Delete("/{taskID}/subtasks/{subtaskID}", handlers.DeleteTaskSubtaskByTaskHandler(db))
}

