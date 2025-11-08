package models

import (
	"time"

	"github.com/Sahil2k07/graphql/internal/enums"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string           `gorm:"not null"`
	Status      enums.TodoStatus `gorm:"default:'PENDING'"`
	UserID      uint             `gorm:"not null"`
	User        User             `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string
	CompletedAt *time.Time
}
