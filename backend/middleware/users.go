package middleware

import (
    "context"
    "net/http"
    "strings"
    "to-do-lister/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        //expect header: Authorization: <token>
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "missing authorization token", http.StatusUnauthorized)
            return
        }

        //allow "Bearer <token>"" or just "<token>"
        parts := strings.Split(token, " ")
        if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
            token = parts[1]
        }

        userID, err := utils.ValidateJWT(token)
        if err != nil {
            http.Error(w, "invalid or expired token", http.StatusUnauthorized)
            return
        }

        //inject userID into context
        ctx := context.WithValue(r.Context(), "userID", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
