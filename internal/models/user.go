package models

import (
	"github.com/Sahil2k07/graphql/internal/enums"
	"gorm.io/gorm"
)

type (
	Profile struct {
		gorm.Model
		UserID    uint
		FirstName string `gorm:"not null"`
		LastName  string `gorm:"not null"`
		Phone     string
		Address   string
		City      string
		State     string
		Country   string
		ZipCode   string
	}

	User struct {
		gorm.Model
		Email    string     `gorm:"not null; uniqueIndex"`
		UserName string     `gorm:"not null"`
		Password string     `gorm:"not null"`
		Role     enums.Role `gorm:"default:'GUEST'"`
		Profile  Profile    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)
