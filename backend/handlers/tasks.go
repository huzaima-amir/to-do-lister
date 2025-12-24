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

func CreateTaskHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		userID := r.Context().Value("userID").(uint)

		var input struct {
			Title           string `json:"title"`
			Description     string `json:"description"`
			Deadline        time.Time `json:"deadline"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		id, err := services.CreateTask(db, input.Title, input.Description, input.Deadline, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(map[string]any{"id": id})
	}
}

func GetTasksHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        var tasks []models.Task

        // Load tasks + subtasks
        if err := db.Preload("SubTasks").
            Where("user_id = ?", userID).
            Order("deadline ASC").
            Find(&tasks).Error; err != nil {
				http.Error(w, "failed to fetch tasks", http.StatusInternalServerError)
				return
        }

        json.NewEncoder(w).Encode(tasks)
    }
}

func GetTaskByIDHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        taskIDStr := chi.URLParam(r, "taskID")
        taskIDUint64, err := strconv.ParseUint(taskIDStr, 10, 64)
        if err != nil {
            http.Error(w, "invalid task ID", http.StatusBadRequest)
            return
        }
        taskID := uint(taskIDUint64)

        var task models.Task
        //ensure the task belongs to the user
        if err := db.Preload("SubTasks").
            Where("id = ? AND user_id = ?", taskID, userID).
            First(&task).Error; err != nil {
				http.Error(w, "task not found", http.StatusNotFound)
				return
        }

        json.NewEncoder(w).Encode(task)
    }
}


func StartTaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskIDStr := chi.URLParam(r, "taskID")
        taskID, _ := strconv.ParseUint(taskIDStr, 10, 64)

        if err := services.StartTask(db, uint(taskID)); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func EndTaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskIDStr := chi.URLParam(r, "taskID")
        taskID, _ := strconv.ParseUint(taskIDStr, 10, 64)

        if err := services.EndTask(db, uint(taskID)); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func DeleteTaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskIDStr := chi.URLParam(r, "taskID")
        taskID, _ := strconv.ParseUint(taskIDStr, 10, 64)

        if err := services.DeleteTask(db, uint(taskID)); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}

func AddSubtasktoTaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskIDStr := chi.URLParam(r, "taskID")
        taskIDUint64, err := strconv.ParseUint(taskIDStr, 10, 64)
        if err != nil {
            http.Error(w, "invalid task ID", http.StatusBadRequest)
            return
        }
        taskID := uint(taskIDUint64)

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

        if err := services.AddSubtaskToTask(db, taskID, input.Title); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]any{
            "message": "subtask created",
        })
    }
}

func DeleteTaskSubtaskByTaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskIDStr := chi.URLParam(r, "taskID")
        subIDStr := chi.URLParam(r, "subtaskID")

        taskID, _ := strconv.ParseUint(taskIDStr, 10, 64)
        subID, _ := strconv.ParseUint(subIDStr, 10, 64)

        if err := services.DeleteTaskSubtaskByTask(db, uint(taskID), uint(subID)); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}

func ToggleSubtaskHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskIDStr := chi.URLParam(r, "taskID")
        subIDStr := chi.URLParam(r, "subtaskID")

        taskID, _ := strconv.ParseUint(taskIDStr, 10, 64)
        subID, _ := strconv.ParseUint(subIDStr, 10, 64)

        var input struct {
            Checked bool `json:"checked"`
        }

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        if err := services.ToggleTaskSubtaskByTask(db, uint(taskID), uint(subID), input.Checked); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}


