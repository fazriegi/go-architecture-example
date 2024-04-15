package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fazriegi/go-architecture-example/config"
	"github.com/fazriegi/go-architecture-example/delivery/http/route"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	godotenv.Load(".env")

	db := config.NewDatabase()
	app := fiber.New()
	port := os.Getenv("PORT")
	route.NewRoute(app, db)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
