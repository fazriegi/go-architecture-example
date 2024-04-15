package route

import (
	"github.com/fazriegi/go-architecture-example/delivery/http"
	"github.com/fazriegi/go-architecture-example/repository"
	"github.com/fazriegi/go-architecture-example/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewUserRoute(app *fiber.App, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := http.NewUserController(userUsecase)

	users := app.Group("/users")
	users.Post("/register", userController.Create)
}
