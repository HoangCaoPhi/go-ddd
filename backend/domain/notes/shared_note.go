package notes

type SharedNote struct {
	SharedID         string `gorm:"type:uuid;primary_key"`
	NoteID           string `gorm:"type:uuid;not null"`
	SharedWithUserID string `gorm:"type:uuid;not null"`
	Permission       string `gorm:"type:enum('read', 'edit');default:'read'"`
}
