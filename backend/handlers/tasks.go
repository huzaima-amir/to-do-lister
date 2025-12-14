package handlers

import (
	"time"
	"gorm.io/gorm"
	"database/models"
  "fmt"
)

func CreateTask(db *gorm.DB, title,description string, deadline time.Time) uint { // creating new task - works
  task := models.Task{Title: title,Description: description, Deadline: deadline, Status: "Pending"}
  db.Create(&task)
  return task.ID
}

func StartTask(db *gorm.DB, taskid uint) error {
    var task models.Task
    if err := db.First(&task, taskid).Error; err != nil {
        return err // task not found
    }

    // Only allow start if task is Pending
    if task.Status == "Finished" {
        return fmt.Errorf("cannot start a finished task")
    }
    if task.Status != "Pending" {
        return fmt.Errorf("task already started")
    }

    task.StartedAt = time.Now()
    task.Status = "In Progress"
    return db.Save(&task).Error
}


func EndTask(db *gorm.DB, taskid uint) error {
    var task models.Task
    if err := db.First(&task, taskid).Error; err != nil {
        return err // if task not found
    }

    //only allow end if task is In Progress
    if task.Status != "In Progress" {
        return fmt.Errorf("cannot end task that hasn't started")
    }

    //mark task as finished
    task.FinishedAt = time.Now()
    task.Status = "Finished"
    if err := db.Save(&task).Error; err != nil {
        return err
    }

    //masscomplete all subtasks for this task
    if err := db.Model(&models.TaskSubTask{}).
        Where("task_id = ?", taskid).
        Update("checked", true).Error; err != nil {
        return err
    }

    return nil
}


func DeleteTask(db *gorm.DB, taskid uint){ 
  db.Delete(&models.Task{}, taskid)
}

func AddSubtaskToTask(db *gorm.DB, pTaskID uint, title string) { // adding subtask to the subtaskchecklist in a specific task
  subTask := models.TaskSubTask{Title:title, Checked: false, TaskID: pTaskID}
  db.Create(&subTask)
}

func DeleteTaskSubtaskByTask(db *gorm.DB, taskID, subtaskID uint) error {
    return db.Where("id = ? AND task_id = ?", subtaskID, taskID). //check parent first?
        Delete(&models.TaskSubTask{}).Error
}

// to mark subtask as checked or unchecked:
func ToggleTaskSubtaskByTask(db *gorm.DB, taskID, subtaskID uint, checked bool) error {
    return db.Model(&models.TaskSubTask{}).
        Where("id = ? AND task_id = ?", subtaskID, taskID).
        Update("checked", checked).Error
}