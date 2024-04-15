package http

import (
	"github.com/fazriegi/go-architecture-example/model"
	"github.com/fazriegi/go-architecture-example/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserController struct {
	UseCase *usecase.UserUsecase
}

func NewUserController(useCase *usecase.UserUsecase) *UserController {
	return &UserController{
		UseCase: useCase,
	}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	user := new(model.User)

	if err := ctx.BodyParser(user); err != nil {
		log.Errorf("Error parsing request body: %s", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing request body",
		})
	}

	status, err := c.UseCase.Create(user)

	if err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(status).JSON(fiber.Map{
		"message": "user created",
	})
}
