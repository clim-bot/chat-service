package models

import (
    "gorm.io/gorm"
)

type Message struct {
    gorm.Model
    Content  string `json:"content"`
    SenderID uint   `json:"sender_id"`
    GroupID  uint   `json:"group_id"`
}
