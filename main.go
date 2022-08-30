package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Print("starting server...")
	app := fiber.New()

	registerRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port: %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
