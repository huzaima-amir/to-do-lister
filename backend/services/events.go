package services

import (
	"time"
	"gorm.io/gorm"
	"to-do-lister/models"
  "fmt"
  )

func CreateEvent(db *gorm.DB, title, description string, startAt, endAt time.Time, location string, online bool, userID uint)(uint, error){ // create new event for specific user
  var user models.User
  if err := db.First(&user, userID).Error; err != nil {
    return 0, fmt.Errorf("user not found")
  }
  event := models.Event{
    Title: title, 
    Description: description, 
    StartsAt: startAt, 
    EndsAt: endAt, 
    Location: location, 
    Online: online, 
    Status: "Upcoming", //default
  }

  if err := db.Create(&event).Error; err != nil { 
      return 0, err   
    } 
  return event.ID, nil
}

func DeleteEvent(db *gorm.DB, eventid uint) error { 
  return db.Delete(&models.Event{}, eventid).Error
}

func AddSubtaskToEvent(db *gorm.DB, eventID uint, title string) error { // adding subtask to the subtaskchecklist in a specific event
    var event models.Event
    if err := db.First(&event, eventID).Error; err != nil { //load the event from DB
        return fmt.Errorf("event not found")
    }

    if event.Status == "Finished" {
        return fmt.Errorf("cannot add subtasks to a finished event")
    }
  subTask := models.EventSubTask{
    Title:title,
     Checked: false, 
     EventID: eventID,
    }
  return db.Create(&subTask).Error 
}

func DeleteEventSubtaskByEvent(db *gorm.DB, eventID, subtaskID uint) error {
    var event models.Event
    if err := db.First(&event, eventID).Error; err != nil { //load the event from DB
        return fmt.Errorf("event not found")
    }
      //prevent deletion or edition of finished event checklist
    if event.Status == "Finished" {
        return fmt.Errorf("cannot delete subtasks from a finished event")
    }
    return db.Where("id = ? AND event_id = ?", subtaskID, eventID). 
        Delete(&models.EventSubTask{}).Error  
}

//To mark subtask as checked/unchecked
func ToggleEventSubtaskByEvent(db *gorm.DB, eventID, subtaskID uint, checked bool) error {
  var event models.Event
  if err := db.First(&event, eventID).Error; err != nil {
    return fmt.Errorf("event not found")
  }
  if event.Status == "Finished" {
    return fmt.Errorf("cannot edit subtasks from a finished event")
  }
  return db.Model(&models.EventSubTask{}).
    Where("id = ? AND event_id = ?", subtaskID, eventID).
    Update("checked", checked).Error 
}

func checkEventStatus(db *gorm.DB) error { // background function that runs in a goroutine forever(for how long the code runs) 
// and updates event status according to passage of time compared to start time and end time
  if err := db.Model(&models.Event{}).
    Where("status = ?", "Upcoming").
    Where("starts_at < ?", time.Now()).
    Update("status", "In Progress").Error; err != nil {
      return err
    }

  if err := db.Model(&models.Event{}).
    Where("status = ?", "In Progress").
    Where("ends_at < ?", time.Now()).
    Update("status", "Finished").Error; err != nil {
      return err
    }

  return nil
}