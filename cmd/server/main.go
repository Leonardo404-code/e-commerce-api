package main

import (
	"log"

	"github.com/Leonargo404-code/e-commerce/internal/database"
	"github.com/Leonargo404-code/e-commerce/internal/env"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func init() {
	database.Connect()
}

func main() {
	app := fiber.New()

	corsSettings := cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "*",
		AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
		ExposeHeaders:    "*, Authorization",
		AllowCredentials: true,
	})

	app.Use(corsSettings)
	app.Static("/images", "./images")

	log.Fatal(app.Listen(":" + env.GetString("PORT")))
}
