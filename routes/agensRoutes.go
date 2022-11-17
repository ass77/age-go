package routes

import (
	"github.com/ass77/age-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func AgensRoutes(route fiber.Router) {
	route.Get("/persons/:relation", controllers.GetPersons)
	route.Get("/person/:relation/:personName", controllers.GetPersonRelation)
	route.Post("/person", controllers.CreatePersonNode)
	route.Post("/connectPersons", controllers.ConnectPersonNode)
	route.Patch("/person/:personName", controllers.UpdatePersonNode)
	route.Delete("/person/:personName", controllers.DeletePersonNode)
	route.Delete("/persons", controllers.DeletePersonsNodes)
}
