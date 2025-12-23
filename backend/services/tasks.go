package services

import (
	"time"
	"gorm.io/gorm"
	"to-do-lister/models"
    "fmt"
)

func CreateTask(db *gorm.DB, title,description string, deadline time.Time, userID uint) (uint, error) { // creating new task for specific user
    var user models.User 
    if err := db.First(&user, userID).Error; err != nil {
        return 0, fmt.Errorf("user not found") 
    }
    task := models.Task{ 
        Title: title, 
        Description: description, 
        Deadline: deadline, 
        Status: "Pending", //default value untill task starts
        Overdue: false, //default value before task gets overdue
        UserID: userID, 
        // StartedAt and FinishedAt will be default zero values
    }
    if err := db.Create(&task).Error; err != nil { 
        return 0, err 
        } 
    return task.ID, nil
}

func StartTask(db *gorm.DB, taskid uint) error {
    var task models.Task
    if err := db.First(&task, taskid).Error; err != nil {
        return fmt.Errorf("task not found") // task not found
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
        return fmt.Errorf("task not found") // if task not found
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

func AddSubtaskToTask(db *gorm.DB, pTaskID uint, title string) error {
    var task models.Task
    if err := db.First(&task, pTaskID).Error; err != nil { //load the task from DB
        return fmt.Errorf("task not found")
    }

    //prevent adding subtasks to finished tasks 
    if task.Status == "Finished" {
        return fmt.Errorf("cannot add subtasks to a finished task")
    }

    //create the subtask
    subTask := models.TaskSubTask{
        Title:  title,
        Checked: false,
        TaskID: pTaskID,
    }
    return db.Create(&subTask).Error
}


func DeleteTaskSubtaskByTask(db *gorm.DB, taskID, subtaskID uint) error {
    var task models.Task
    if err := db.First(&task, taskID).Error; err != nil {
        return fmt.Errorf("task not found")
    }
    // prevent edition of finished task's subtasks
    if task.Status == "Finished" {
        return fmt.Errorf("cannot delete subtasks from a finished task")
    }
    return db.Where("id = ? AND task_id = ?", subtaskID, taskID). 
        Delete(&models.TaskSubTask{}).Error 
}

// to mark subtask as checked or unchecked:
func ToggleTaskSubtaskByTask(db *gorm.DB, taskID, subtaskID uint, checked bool) error {
    var task models.Task
    if err := db.First(&task, taskID).Error; err != nil {
        return fmt.Errorf("task not found")
    }
    // prevent edition of finished task's subtasks
    if task.Status == "Finished" {
// only displaying error for unchecking, because finishing a task automatically masschecks all subtasks, 
// so there are none left to check
        return fmt.Errorf("cannot uncheck subtasks from a finished task") 
    }
    return db.Model(&models.TaskSubTask{}).
        Where("id = ? AND task_id = ?", subtaskID, taskID).
        Update("checked", checked).Error 
}

// Need to add new goroutine and function to update task overdue status 
// if task deadline passes before task finished-  

func UpdateTaskOverdueStatus(db *gorm.DB) error{// background function that runs in a goroutine for"ever"(for how long the code runs) 
// and marks task as overdue if applicable
// putting the logic in this function, but will initialize the goroutine with the forloop in the main function (database intialization issues)
    return db.Model(&models.Task{}). 
        Where("status != ?", "Finished").
        Where("deadline < ?", time.Now()).  //comparison operator for time.time only works in sql operations!!
        Update("overdue", true).Error
}

