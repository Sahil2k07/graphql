package main

import (
	"github.com/Sahil2k07/graphql/internal/configs"
	"github.com/Sahil2k07/graphql/internal/database"
	"github.com/Sahil2k07/graphql/internal/models"
	"github.com/charmbracelet/log"
)

func main() {
	configs.LoadConfig()
	database.Connect()

	// Migration - 1
	models := []any{&models.User{}, &models.Profile{}, &models.Todo{}}

	err := database.DB.AutoMigrate(models...)
	if err != nil {
		log.Errorf("Migration failed: %v", err)
		return
	}

	log.Infof("Migration Completed Successfully!")
}
