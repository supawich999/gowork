package routes

import (
	"gowork/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("api/v1/list")
	}) // /api/v1/list

	app.Post("/login", controllers.TestLogin)
	app.Get("/dog", controllers.GetDogs)
	app.Post("/dog", controllers.AddDog)
	app.Put("/dog/:id", controllers.UpdateDog)
	app.Delete("/dog/:id", controllers.RemoveDog)
	app.Get("/dog/v2", controllers.GetDog)

	v1.Get("/user", controllers.GetUsers)
	v1.Get("/userbyid", controllers.GetUser)
	v1.Post("/user", controllers.AddUser)
	v1.Put("/user/:id", controllers.UpdateUser)
	v1.Delete("/user/:id", controllers.RemoveUser)

}
