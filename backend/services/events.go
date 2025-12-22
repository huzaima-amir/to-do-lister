package services

import (
	"time"
	"gorm.io/gorm"
	"to-do-lister/models"
  )

func CreateEvent(db *gorm.DB, title, description string, startAt, endAt time.Time, location string, online bool) uint { // create new event
  event := models.Event{Title: title, Description: description, StartsAt: startAt, EndsAt: endAt, Location: location, Online: online, Status: "Upcoming"}
  db.Create(&event)
  return event.ID  // consider user ID !!!TODO
}

func DeleteEvent(db *gorm.DB, eventid uint){ 
  db.Delete(&models.Event{}, eventid)
}

func AddSubtaskToEvent(db *gorm.DB, pEventID uint, title string) { // adding subtask to the subtaskchecklist in a specific event
  subTask := models.EventSubTask{Title:title, Checked: false, EventID: pEventID}
  db.Create(&subTask)  // can only work if event isnt finished yet !!!TODO
}

func DeleteEventSubtaskByEvent(db *gorm.DB, eventID, subtaskID uint) error {
    return db.Where("id = ? AND event_id = ?", subtaskID, eventID). 
        Delete(&models.EventSubTask{}).Error  // can only work if event hasnt ended yet !!!TODO
}


//To mark subtask as checked/unchecked
func ToggleEventSubtaskByEvent(db *gorm.DB, eventID, subtaskID uint, checked bool) error {
    return db.Model(&models.EventSubTask{}).
        Where("id = ? AND event_id = ?", subtaskID, eventID).
        Update("checked", checked).Error // can only work if event hasnt ended yet !!!TODO
}

//  need a new function goroutine  that checks in parallel if event has started and ended to update status accordingly !!!TODO

func checkEventStatus(){ // background function that runs in a goroutine forever(for how long the code runs) 
// and updates event status according to passage of time compared to start time and end time

}