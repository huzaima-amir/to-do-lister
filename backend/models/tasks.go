package models

import (
	"time"
	)

type Task struct {
  ID          uint   `gorm:"primaryKey"`
  UserID      uint `gorm:"not null"`
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

