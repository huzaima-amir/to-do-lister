package handlers

import (
	"database/models"
	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, title, desc string) uint {
	tag := models.Tag{Title: title, Description: desc}
	db.Select("ID", "Title", "Description").Create(&tag)
	return tag.ID
}

func DeleteTag (db *gorm.DB, tagID uint){ 
	db.Model(&models.Task{}).Association("Tags").Clear()  
	db.Model(&models.Event{}).Association("Tags").Clear()
	db.Delete(&models.Tag{}, tagID)
}


func AddTagToTask(db *gorm.DB, taskID, tagID uint) {
    var task models.Task
    var tag models.Tag
    db.First(&task, taskID)
    db.First(&tag, tagID)
    db.Model(&task).Association("Tags").Append(&tag)
}


func RemoveTagFromTask(db *gorm.DB, taskID, tagID uint) {
    var task models.Task
    var tag models.Tag
    db.First(&task, taskID)
    db.First(&tag, tagID)
    db.Model(&task).Association("Tags").Delete(&tag)
}


func AddTagToEvent(db *gorm.DB, eventID, tagID uint) {
    var event models.Event
    var tag models.Tag
    db.First(&event, eventID)
    db.First(&tag, tagID)
    db.Model(&event).Association("Tags").Append(&tag)
}


func RemoveTagFromEvent(db *gorm.DB, eventID, tagID uint) {
    var event models.Event
    var tag models.Tag
    db.First(&event, eventID)
    db.First(&tag, tagID)
    db.Model(&event).Association("Tags").Delete(&tag)
}

