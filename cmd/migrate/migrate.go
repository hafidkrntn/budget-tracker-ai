package main

import (
	"backend-go/config"
	"backend-go/internal/model/migrate"
	"log"
)

func main() {
	config.InitDB() // ✅ menggunakan koneksi global
	db := config.DbConn

	log.Println("🔧 Running AutoMigrate...")

	if err := db.AutoMigrate(
		&migrate.User{},
		&migrate.LogError{},
		&migrate.Category{},
		&migrate.Transaction{},
	); err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	log.Println("✅ Migration completed successfully!")
}
