package main

import (
	"github.com/aydogduyusuf/Network-Device-Monitoring/web"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	web.SetupRoutes(app)

	port := ":8080"
	log.Println("Server listening on port", port)
	log.Fatal(app.Listen(port))
}
