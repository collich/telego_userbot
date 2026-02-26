package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type DB struct {
	conn *sql.DB
}


func New(dbPath string) (*DB, error) {
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db := &DB{conn: conn}

	if err := db.initSchema(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to initialize schema: %w", err)
	}

	return db, nil
}

func (db *DB) initSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS downloads (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		creator_name TEXT NOT NULL,
		description TEXT,
		filename TEXT NOT NULL,
		file_path TEXT NOT NULL,
		message_id INTEGER NOT NULL,
		group_id INTEGER NOT NULL,
		downloaded_at DATETIME NOT NULL,
		UNIQUE(group_id, filename)
	);

	CREATE INDEX IF NOT EXISTS idx_creator_name ON downloads(creator_name);
	CREATE INDEX IF NOT EXISTS idx_message_id ON downloads(message_id);
	CREATE INDEX IF NOT EXISTS idx_group_id ON downloads(group_id);
	`

	_, err := db.conn.Exec(schema)
	return err
}

func (db *DB) SaveDownload(download *Download) error {
	query := `
		INSERT INTO downloads (creator_name, description, filename, file_path, message_id, group_id, downloaded_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := db.conn.Exec(
		query,
		download.CreatorName,
		download.Description,
		download.Filename,
		download.FilePath,
		download.MessageID,
		download.GroupID,
		download.DownloadedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save download: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	download.ID = int(id)
	return nil
}

func (db *DB) GetDownloadsByGroupID(groupID int64) ([]Download, error) {
	query := `SELECT id, creator_name, description, filename, file_path, message_id, group_id, downloaded_at
	          FROM downloads WHERE group_id = ?`

	rows, err := db.conn.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var downloads []Download
	for rows.Next() {
		var d Download
		err := rows.Scan(&d.ID, &d.CreatorName, &d.Description, &d.Filename, &d.FilePath, &d.MessageID, &d.GroupID, &d.DownloadedAt)
		if err != nil {
			return nil, err
		}
		downloads = append(downloads, d)
	}

	return downloads, rows.Err()
}

func (db *DB) GroupExists(groupID int64) (bool, error) {
	var count int
	err := db.conn.QueryRow("SELECT COUNT(*) FROM downloads WHERE group_id = ?", groupID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func NewDownload(creatorName, description, filename, filePath string, messageID, groupID int64) *Download {
	return &Download{
		CreatorName:  creatorName,
		Description:  description,
		Filename:     filename,
		FilePath:     filePath,
		MessageID:    messageID,
		GroupID:      groupID,
		DownloadedAt: time.Now(),
	}
}
