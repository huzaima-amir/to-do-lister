package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"to-do-lister/services"

	//"to-do-lister/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)
//  TODO!!:add delete tags, and get all tags
//  to maybe display tags list or in order to get list of tags to choose from to add to tasks and events

func CreateTagHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        var input struct {
            Title       string `json:"title"`
            Description string `json:"description"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        id, err := services.CreateTag(db, input.Title, input.Description, userID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]any{"id": id})
    }
}


func AddTagToTaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        taskIDStr := chi.URLParam(r, "taskID")
        taskIDUint64, err := strconv.ParseUint(taskIDStr, 10, 64)
        if err != nil {
            http.Error(w, "invalid task ID", http.StatusBadRequest)
            return
        }
        taskID := uint(taskIDUint64)

        var input struct {
            TagID uint `json:"tag_id"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if input.TagID == 0 {
            http.Error(w, "tag_id is required", http.StatusBadRequest)
            return
        }

        if err := services.AddTagToTask(db, taskID, input.TagID, userID); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]any{
            "message": "tag added to task",
        })
    }
}


func RemoveTagFromTaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        taskIDStr := chi.URLParam(r, "taskID")
        taskIDUint64, err := strconv.ParseUint(taskIDStr, 10, 64)
        if err != nil {
            http.Error(w, "invalid task ID", http.StatusBadRequest)
            return
        }
        taskID := uint(taskIDUint64)

        var input struct {
            TagID uint `json:"tag_id"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if input.TagID == 0 {
            http.Error(w, "tag_id is required", http.StatusBadRequest)
            return
        }

        if err := services.RemoveTagFromTask(db, taskID, input.TagID, userID); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}


func AddTagToEventHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        eventIDStr := chi.URLParam(r, "eventID")
        eventIDUint64, err := strconv.ParseUint(eventIDStr, 10, 64)
        if err != nil {
            http.Error(w, "invalid event ID", http.StatusBadRequest)
            return
        }
        eventID := uint(eventIDUint64)

        var input struct {
            TagID uint `json:"tag_id"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if input.TagID == 0 {
            http.Error(w, "tag_id is required", http.StatusBadRequest)
            return
        }

        if err := services.AddTagToEvent(db, eventID, input.TagID, userID); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]any{
            "message": "tag added to event",
        })
    }
}

func RemoveTagFromEventHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        eventIDStr := chi.URLParam(r, "eventID")
        eventIDUint64, err := strconv.ParseUint(eventIDStr, 10, 64)
        if err != nil {
            http.Error(w, "invalid event ID", http.StatusBadRequest)
            return
        }
        eventID := uint(eventIDUint64)

        var input struct {
            TagID uint `json:"tag_id"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if input.TagID == 0 {
            http.Error(w, "tag_id is required", http.StatusBadRequest)
            return
        }

        if err := services.RemoveTagFromEvent(db, eventID, input.TagID, userID); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}
