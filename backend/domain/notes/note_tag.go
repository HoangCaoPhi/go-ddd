package notes

type NoteTag struct {
	NoteID string `gorm:"type:uuid;primary_key"`
	TagID  string `gorm:"type:uuid;primary_key"`
}
