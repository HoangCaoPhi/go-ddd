package notes

import "time"

type Reminder struct {
	ReminderID   string    `gorm:"type:uuid;primary_key"`
	NoteID       string    `gorm:"type:uuid;not null"`
	ReminderTime time.Time `gorm:"not null"`
	IsActive     bool      `gorm:"default:true"`
}
