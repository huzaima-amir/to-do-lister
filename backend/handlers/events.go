package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"to-do-lister/models"
	"to-do-lister/services"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateEventHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		userID := r.Context().Value("userID").(uint)

		var input struct {
			Title   	string `json:"title"`
			Description	string `json:"description"`
			StartsAt	time.Time `json:"starts_at"`
			EndsAt		time.Time `json:"ends_at"`
			Location	string `json:"location"`
			Online		bool `json:"online"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		
		id, err := services.CreateEvent(db, 
			input.Title, 
			input.Description, 
			input.StartsAt, 
			input.EndsAt,
			input.Location, 
			input.Online,  
			userID) 
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			}

		json.NewEncoder(w).Encode(map[string]any{"id": id})
	}
}

func GetEventsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		userID := r.Context().Value("userID").(uint)

		var events []models.Event

		if err := db.Preload("SubTasks").
			Where("user_id = ?", userID).
			Order("starts_at ASC").
			Find(&events).Error; err != nil {
				http.Error(w, "failed to fetch events", http.StatusInternalServerError)
				return
			}

		json.NewEncoder(w).Encode(events)
	}
}

func GetEventByIDHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(uint)

		eventIDStr := chi.URLParam(r, "taskID")
		eventIDUint64, err := strconv.ParseUint(eventIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid event ID", http.StatusBadRequest)
			return
		}
		eventID := uint(eventIDUint64)

		var event models.Event
		if err := db.Preload("SubTasks").
		Where("id = ? AND user_id = ?", eventID, userID).
		First(&event).Error; err != nil {
			http.Error(w, "event not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(event)
	}
}

// func UpdateEventStatusHandler(db *gorm.DB) http.HandlerFunc {  }  TODO!!!



func DeleteEventHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventIDStr := chi.URLParam(r, "eventID")
		eventID, _ := strconv.ParseUint(eventIDStr, 10, 64)

		if err := services.DeleteEvent(db, uint(eventID)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func AddSubtasktoEventHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventIDStr := chi.URLParam(r, "eventID")
		eventIDUint64, err := strconv.ParseUint(eventIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid event ID", http.StatusBadRequest)
			return
		}
		eventID := uint(eventIDUint64)

		var input struct {
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if input.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}

		if err := services.AddSubtaskToEvent(db, eventID, input.Title); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"message" : "subtask created",
		})
	}
}


func DeleteSubTaskByEventHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventIDStr := chi.URLParam(r, "eventID")
		subIDStr := chi.URLParam(r, "subtaskID")

		eventID, _ := strconv.ParseUint(eventIDStr, 10, 64)
		subID, _ := strconv.ParseUint(subIDStr, 10, 64)

		if err := services.DeleteEventSubtaskByEvent(db, uint(eventID), uint(subID)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ToggleEventSubtaskHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventIDStr := chi.URLParam(r, "eventID")
		subIDStr  := chi.URLParam(r, "subtaskID")

		eventID, _ := strconv.ParseUint(eventIDStr, 10, 64)
		subID, _ := strconv.ParseUint(subIDStr, 10, 64)

		var input struct {
			Checked bool `json: "checked"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if err := services.ToggleEventSubtaskByEvent(db, uint(eventID), uint(subID), input.Checked); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}