package tags

import "github.com/google/uuid"

type Tag struct {
	TagID  uuid.UUID `gorm:"type:uniqueidentifier;default:NEWID()"`
	UserID uuid.UUID `gorm:"type:uniqueidentifier;not null"`
	Name   string    `gorm:"type:varchar(255);not null"`
}

func CreateTag(name string, userId uuid.UUID) *Tag {
	return &Tag{uuid.New(), userId, name}
}
