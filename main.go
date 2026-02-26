package main

import (
	"log"
	"os"
	"tellego_userbot/config"
	"tellego_userbot/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("🚀 Starting Telegram Userbot...")
	log.Printf("📱 Phone: %s", cfg.TelegramPhone)
	log.Printf("📂 Download Dir: %s", cfg.DownloadDir)
	log.Printf("💾 Database: %s", cfg.DatabasePath)
	log.Printf("🎯 Target Group: %s", cfg.TargetGroupName)
	log.Printf("📦 Session Dir: %s", cfg.SessionDir)

	db, err := database.New(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	log.Println("✓ Database initialized")

	if err := os.MkdirAll(cfg.DownloadDir, 0755); err != nil {
		log.Fatalf("Failed to create download directory: %v", err)
	}
	if err := os.MkdirAll(cfg.SessionDir, 0755); err != nil {
		log.Fatalf("Failed to create session directory: %v", err)
	}
}