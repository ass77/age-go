package routes

import (
	"github.com/ass77/age-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func AgensRoutes(route fiber.Router) {
	route.Get("/persons", controllers.GetPersons)
	route.Get("/person/:personId", controllers.GetPerson)
	route.Post("/person", controllers.CreatePerson)
	route.Post("/connectPersons", controllers.ConnectPersons)
	route.Patch("/person/:personId", controllers.UpdatePerson)
	route.Delete("/person/:personName", controllers.DeletePerson)
	route.Delete("/persons", controllers.DeletePersons)
}
