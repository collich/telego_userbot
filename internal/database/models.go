package database

import "time"

type Download struct {
	ID           int       `db:"id"`
	CreatorName  string    `db:"creator_name"`
	Description  string    `db:"description"`
	Filename     string    `db:"filename"`
	FilePath     string    `db:"file_path"`
	MessageID    int64     `db:"message_id"`
	GroupID      int64     `db:"group_id"` // Album GroupedID or message_id if solo
	DownloadedAt time.Time `db:"downloaded_at"`
}

type ParsedMessage struct {
	CreatorName string
	Description string
	IsValid     bool
}

type MediaFile struct {
	Type     string
	Filename string
	FilePath string
}
