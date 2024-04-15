package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRoute(app *fiber.App, db *gorm.DB) {
	NewUserRoute(app, db)
}
