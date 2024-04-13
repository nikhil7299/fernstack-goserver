package main

import (
	"log"
	"os"
	"server/package/configs"
	"server/package/routes"
	mongoDB "server/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	// used to autoload .env - "_" is used to avoid error of package not being used
	// _"github.com/joho/godotenv/autoload"
	// "github.com/go-playground/validator/v10"
)

func main() {
	godotenv.Load()
	config := configs.FiberConfig()
	app := fiber.New(config)

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "[Fiber] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "02-01-2006 03:04:05 PM",
		TimeZone:   "Local",
	}))

	err := mongoDB.MongoConnect()
	if err != nil {
		panic(err)
	}
	defer mongoDB.MongoDisconnect()

	routes.AuthRoutes(app)

	port := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen(":" + port))

}
