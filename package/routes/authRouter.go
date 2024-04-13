package routes

import (
	"server/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Get("/", controllers.GetUser)

	authRouter := app.Group("/auth")

	authRouter.Post("/signUpEmail", controllers.SignUpEmail)
	authRouter.Post("/logInEmail", controllers.LogInEmail)
	authRouter.Post("/logInGoogle", controllers.LogInGoogle)

}
