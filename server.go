package main

import (
	"log"
	"os"

	"github.com/durgaprasad97005/fiberCrudBooks/src"
)

// main function where the execution starts
func main() {
	app := src.SetupApp()

	port := os.Getenv("PORT")
	// fallback
	if port == "" {
		port = "3000"
	}

	log.Println("The server is running at port:" + port)
	log.Fatal(app.Listen(":" + port))
}