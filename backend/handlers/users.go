package handlers

import (
"encoding/json" 
"net/http" 
//"strconv" 
// "github.com/go-chi/chi/v5" 
"gorm.io/gorm" 
"to-do-lister/services" 
"to-do-lister/utils"
)

func SignUpHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input struct {
            Name     string `json:"name"`
            UserName string `json:"username"`
            Password string `json:"password"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        id, err := services.CreateUser(db, input.Name, input.UserName, input.Password)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        json.NewEncoder(w).Encode(map[string]any{"id": id})
    }
}


func LogInHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input struct {
            UserName string `json:"username"`
            Password string `json:"password"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        userID, err := services.ValidateUserCredentials(db, input.UserName, input.Password)
        if err != nil {
            http.Error(w, "invalid username or password", http.StatusUnauthorized)
            return
        }

        token, err := utils.GenerateJWT(userID)
        if err != nil {
            http.Error(w, "failed to generate token", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(map[string]any{"token": token})
    }
}


func LogOutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "logged out",
		})
	}
}

func DeleteUserHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        if err := services.DeleteUser(db, userID); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}


func ChangeNameHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        var input struct {
            NewName string `json:"new_name"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if err := services.ChangeName(db, userID, input.NewName); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}


func ChangeUsernameHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        var input struct {
            NewUsername string `json:"new_username"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if err := services.ChangeUsername(db, userID, input.NewUsername); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}


func ChangePasswordHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        var input struct {
            OldPassword string `json:"old_password"`
            NewPassword string `json:"new_password"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if err := services.ChangePassword(db, userID, input.OldPassword, input.NewPassword); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}
