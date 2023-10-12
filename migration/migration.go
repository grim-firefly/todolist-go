package migration

import (
	"github.com/grim-firefly/todolist-go/database"
	"github.com/grim-firefly/todolist-go/models"
)

func Migrate() {
	db := database.GetDB()

	// todo migrating
	db.AutoMigrate(&models.Todo{})
}
