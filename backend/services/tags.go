package services

import (
	"to-do-lister/models"
	"gorm.io/gorm"
    "fmt"
)

func CreateTag(db *gorm.DB, title, desc string, userID uint) (uint, error) {
    tag := models.Tag{
        Title:       title,
        Description: desc,
        UserID:      userID,
    }
    if err := db.Create(&tag).Error; err != nil {
        return 0, err
    }

    return tag.ID, nil
}

func DeleteTag(db *gorm.DB, tagID uint, userID uint) error {
    var tag models.Tag
    // ensure tag exists and belongs to the specific user
    if err := db.Where("id = ? AND user_id = ?", tagID, userID).First(&tag).Error; err != nil {
        return fmt.Errorf("tag not found")
    }

    if err := db.Model(&models.Task{}).   // remove tag from user's tasks
        Where("user_id = ?", userID).
        Association("Tags").
        Delete(&tag); err != nil {
        return err
    }

    if err := db.Model(&models.Event{}).  // remove tag from user's events
        Where("user_id = ?", userID).
        Association("Tags").
        Delete(&tag); err != nil {
        return err
    }

    return db.Delete(&tag).Error
}



func AddTagToTask(db *gorm.DB, taskID, tagID, userID uint) error {
    var task models.Task
    var tag models.Tag

    //load task and ensure ownership of specific user
    if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
        return fmt.Errorf("task not found")
    }
    // prevent tagging finished tasks
    if task.Status == "Finished" {
        return fmt.Errorf("cannot tag a finished task")
    }
    // load tag and ensure ownership
    if err := db.Where("id = ? AND user_id = ?", tagID, userID).First(&tag).Error; err != nil {
        return fmt.Errorf("tag not found")
    }

    return db.Model(&task).Association("Tags").Append(&tag)
}



func RemoveTagFromTask(db *gorm.DB, taskID, tagID, userID uint) error {
    var task models.Task
    var tag models.Tag

    if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
        return fmt.Errorf("task not found")
    }
    if task.Status == "Finished" {
        return fmt.Errorf("cannot remove tags from a finished task")
    }
    if err := db.Where("id = ? AND user_id = ?", tagID, userID).First(&tag).Error; err != nil {
        return fmt.Errorf("tag not found")
    }

    return db.Model(&task).Association("Tags").Delete(&tag)
}


func AddTagToEvent(db *gorm.DB, eventID, tagID, userID uint) error {
    var event models.Event
    var tag models.Tag

    if err := db.Where("id = ? AND user_id = ?", eventID, userID).First(&event).Error; err != nil {
        return fmt.Errorf("event not found")
    }
    if event.Status == "Finished" {
        return fmt.Errorf("cannot tag a finished event")
    }
    if err := db.Where("id = ? AND user_id = ?", tagID, userID).First(&tag).Error; err != nil {
        return fmt.Errorf("tag not found")
    }

    return db.Model(&event).Association("Tags").Append(&tag)
}



func RemoveTagFromEvent(db *gorm.DB, eventID, tagID, userID uint) error {
    var event models.Event
    var tag models.Tag

    if err := db.Where("id = ? AND user_id = ?", eventID, userID).First(&event).Error; err != nil {
        return fmt.Errorf("event not found")
    }
    if event.Status == "Finished" {
        return fmt.Errorf("cannot remove tags from a finished event")
    }
    if err := db.Where("id = ? AND user_id = ?", tagID, userID).First(&tag).Error; err != nil {
        return fmt.Errorf("tag not found")
    }

    return db.Model(&event).Association("Tags").Delete(&tag)
}


