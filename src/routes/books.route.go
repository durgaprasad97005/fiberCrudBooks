package routes

import (
	"github.com/durgaprasad97005/fiberCrudBooks/src/controllers"
	"github.com/gofiber/fiber/v3"
)

// function that create routes for CRUD operations on books
func BooksRoutes(app *fiber.App) {
	books := app.Group("/books")

	books.Get("/", controllers.GetBooks)
	books.Get("/:id", controllers.GetBookById)
	books.Post("/", controllers.CreateBook)
	books.Put("/:id", controllers.UpdateBookById)
	books.Delete("/:id", controllers.DeleteBookById)
}