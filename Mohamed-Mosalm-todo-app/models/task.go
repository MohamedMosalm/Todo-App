package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
