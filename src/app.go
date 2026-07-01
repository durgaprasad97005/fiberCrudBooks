package src

import (
	"github.com/durgaprasad97005/fiberCrudBooks/src/db"
	"github.com/durgaprasad97005/fiberCrudBooks/src/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	_ = godotenv.Load()

	db.ConnectDB()

	routes.BooksRoutes(app)

	return app
}