package handlers

import (
	"time"
	"gorm.io/gorm"
	"database/models"
  )

func CreateEvent(db *gorm.DB, title, description string, startAt, endAt time.Time, location string, online bool) uint { // create new event
  event := models.Event{Title: title, Description: description, StartsAt: startAt, EndsAt: endAt, Location: location, Online: online, Status: "Upcoming"}
  db.Create(&event)
  return event.ID
}

func DeleteEvent(db *gorm.DB, eventid uint){ // remove task = works
  db.Delete(&models.Event{}, eventid)
}

func AddSubtaskToEvent(db *gorm.DB, pEventID uint, title string) { // adding subtask to the subtaskchecklist in a specific event
  subTask := models.EventSubTask{Title:title, Checked: false, EventID: pEventID}
  db.Create(&subTask)
}

func DeleteEventSubtaskByEvent(db *gorm.DB, eventID, subtaskID uint) error {
    return db.Where("id = ? AND event_id = ?", subtaskID, eventID). 
        Delete(&models.EventSubTask{}).Error
}


//To mark subtask as checked/unchecked
func ToggleEventSubtaskByEvent(db *gorm.DB, eventID, subtaskID uint, checked bool) error {
    return db.Model(&models.EventSubTask{}).
        Where("id = ? AND event_id = ?", subtaskID, eventID).
        Update("checked", checked).Error
}