package routes

import (
    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
    "to-do-lister/handlers"
    "to-do-lister/middleware"
)

func UserRoutes(r chi.Router, db *gorm.DB) {

    //public endpoints for signup and login
    r.Post("/signup", handlers.SignUpHandler(db))
    r.Post("/login", handlers.LoginHandler(db))

    //protected endpoints for user detail edits and 
    r.Group(func(pr chi.Router) {
        pr.Use(middleware.AuthMiddleware)

        pr.Put("/name", handlers.ChangeNameHandler(db))
        pr.Put("/password", handlers.ChangePasswordHandler(db))
        pr.Put("/username", handlers.ChangeUsernameHandler(db))
        pr.Delete("/delete", handlers.DeleteUserHandler(db))
    })
}
