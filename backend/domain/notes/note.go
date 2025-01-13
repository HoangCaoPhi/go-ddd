package notes

import (
	"time"
)

type Note struct {
	NoteID      string       `gorm:"type:uuid;primary_key"`
	UserID      string       `gorm:"type:uuid;not null"`
	Title       string       `gorm:"type:varchar(255)"`
	Content     string       `gorm:"type:text"`
	IsChecklist bool         `gorm:"default:false"`
	Status      string       `gorm:"type:enum('active', 'completed', 'archived');default:'active'"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
	Reminders   []Reminder   `gorm:"foreignKey:NoteID"`
	SharedNotes []SharedNote `gorm:"foreignKey:NoteID"`
}
