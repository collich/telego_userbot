package main

import (
	"log"
	"tellego_userbot/config"
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
}