package models

import (
	"time"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name string
	UserName string   `gorm:"unique;not null"` // unique key can be edited?
	Password string  `gorm:"not null"`
	CreatedAt time.Time // easier auditing?
	UpdatedAt time.Time
	Tasks  []Task  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Events []Event `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"` 
	//one to many association for tasks and events + All task and event data including tags gets deleted when user account is deleted.
	Tags   []Tag   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}