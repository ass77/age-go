package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ass77/age-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupRoutes(app *fiber.App) {

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":   true,
			"message":   "pong",
			"timestamp": time.Now().Unix(),
		})
	})

	api := app.Group("/")
	routes.AgensRoutes(api.Group("/agens"))

}

// TODO generic graphDB -> use metadata to create vertexes and match edges
func main() {

	var ENV_TYPE = os.Getenv("APP_ENV")

	if ENV_TYPE == "local" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")

		}

	}
	fmt.Println("ENV_TYPE: ", ENV_TYPE)

	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE, PATCH",
		AllowCredentials: true,
		ExposeHeaders:    "Origin",
	}))

	setupRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
