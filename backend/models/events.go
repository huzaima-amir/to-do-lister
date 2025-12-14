package models

import (
	"time"
)

type Event struct {
  ID          uint   `gorm:"primaryKey"`
  Title       string
  Description string
  StartsAt    time.Time
  EndsAt      time.Time
  Location    string
  Online      bool
  Status      string
  SubTaskCheckList []EventSubTask `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;"`
  Tags        []Tag `gorm:"many2many:event_tags;"`
}

// All tasks and events have checklist that has multiple subTasks
type EventSubTask struct { 
  ID uint `gorm:"primaryKey"`
  Title string  
  Checked bool //finished subtasks can be checked off in checklists
  EventID uint `gorm:"constraint:OnDelete:CASCADE;"`
}

