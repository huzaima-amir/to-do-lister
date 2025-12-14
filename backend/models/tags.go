package models

type Tag struct { // Many-to-Many, cause Events and Tasks can use multiple tags, 
// even shared tags, so tags can be used by multiple tasks or events as well.
  ID          uint   `gorm:"primaryKey"`
  Title       string
  Description string
  Events      []Event `gorm:"many2many:event_tags;"`
  Tasks       []Task  `gorm:"many2many:task_tags;"`
}
