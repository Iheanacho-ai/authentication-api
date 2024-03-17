package initializers

import (
	"github.com/Iheanacho-ai/authentication-api.git/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}