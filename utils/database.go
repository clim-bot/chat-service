package utils

import (
    "database/sql"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/clim-bot/chat-service/models"
)

func SetupDatabase() (*gorm.DB, *sql.DB) {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    db.AutoMigrate(&models.User{}, &models.Message{})

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("Failed to get SQL DB from GORM DB:", err)
    }

    return db, sqlDB
}
