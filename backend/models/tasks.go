package models

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
  ID          uint   `gorm:"primaryKey"`
  Title       string
  Description string
  SubTaskCheckList []TaskSubTask `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;"`
  Status      string
  Deadline    time.Time
  StartedAt   time.Time
  FinishedAt  time.Time
  Overdue     bool
  Tags        []Tag `gorm:"many2many:task_tags;"`
}

// All tasks and events have checklist that has multiple subTasks only associated with the specific task
type TaskSubTask struct {
  ID uint `gorm:"primaryKey"`
  Title string 
  Checked bool
  TaskID uint  `gorm:"constraint:OnDelete:CASCADE;"`
}


// methods:
func (t *Task) AfterCreate(db *gorm.DB) (err error) {
  if t.Deadline.Before(time.Now()) && t.Status != "Finished" { // check for deadline passing before task is finished to trigger Overdue warning/
    db.Model(&Task{}).Where("id = ?", t.ID).Update("overdue", true)
    return db.Save(&t).Error
  }
  return
}