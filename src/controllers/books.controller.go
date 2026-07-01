package controllers

import (
	"github.com/durgaprasad97005/fiberCrudBooks/src/db"
	"github.com/durgaprasad97005/fiberCrudBooks/src/models"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller to get all books
func GetBooks(c fiber.Ctx) error {
	// Get collection from db
	coll := db.GetCollection("books")

	if coll == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error", 
		})
	}

	// Return all books from collection
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error: " + err.Error(), 
		})
	}

	// Parse cursor to books slice
	var books []models.Book
	if err := cursor.All(c.Context(), &books); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error: " + err.Error(),
		})
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully returned the books from database", 
		"books": books,
	})
}

// Controller to get a book by its ID
func GetBookById(c fiber.Ctx) error {
	// Get the collection
	coll := db.GetCollection("books")

	// Get object id from query parameters of the request
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not a valid Id", 
		})
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id", 
		})
	}

	// Find the document in collection
	var book models.Book

	err = coll.FindOne(c.Context(), bson.M{"_id": objId}).Decode(&book)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error finding book: " + err.Error(), 
		})
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully returned requested document", 
		"book": book, 
	})
}

// Controller to create / post a book
func CreateBook(c fiber.Ctx) error {
	// Getting the collection
	coll := db.GetCollection("books")
	if coll == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error", 
		})
	}

	// Parse request body to get book info
	var book models.Book
	if err := c.Bind().Body(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing the body: " + err.Error(), 
		})
	}

	// Creating the book
	result, err := coll.InsertOne(c.Context(), book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating book: " + err.Error(), 
		})
	}

	// Success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created", 
		"result": result, // check whether complete book info is coming or only id is coming
	})
}

// Controller to update a book
func UpdateBookById(c fiber.Ctx) error {
	
	// Get the collection
	coll := db.GetCollection("books")
	if coll == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error", 
		})
	}

	// Get id query parameter from request
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not a valid Id", 
		})
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id", 
		})
	}

	// Check whether the book exists or not in collection
	var book models.Book
	err = coll.FindOne(c.Context(), bson.M{"_id": objId}).Decode(&book)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error finding id: " + err.Error(), 
		})
	}

	// Parse the request body
	var body models.Book
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing body: " + err.Error(), 
		})
	}

	// Update the book
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objId}, bson.M{"$set": body})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating book: " + err.Error(), 
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to match document for update", 
		})
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Updated successfully", 
		"result": result,
	})
}

// Controller to delete a book
func DeleteBookById(c fiber.Ctx) error {
	// Get the collection
	coll := db.GetCollection("books")
	if coll == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error", 
		})
	}

	// Get id query parameter from request
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not a valid Id", 
		})
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id", 
		})
	}

	// Check whether the book exists or not in collection
	var book models.Book
	err = coll.FindOne(c.Context(), bson.M{"_id": objId}).Decode(&book)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error finding id: " + err.Error(), 
		})
	}

	// Deleting the document from collection
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting book: " + err.Error(), 
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "There isn't any document to delete", 
		})
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted successfully",
		"result": result,
	})
}