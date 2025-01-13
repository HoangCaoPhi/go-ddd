package entities

import (
	"phihc116/go-task/backend/domain/notes"
	"time"
)

type User struct {
	UserID       string       `gorm:"type:uuid;primary_key"`
	Email        string       `gorm:"type:varchar(255);unique;not null"`
	PasswordHash string       `gorm:"type:varchar(255);not null"`
	FullName     string       `gorm:"type:varchar(255)"`
	CreatedAt    time.Time    `gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `gorm:"autoUpdateTime"`
	Notes        []notes.Note `gorm:"foreignKey:UserID"`
}
