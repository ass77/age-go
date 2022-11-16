package routes

import (
	"github.com/ass77/age-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func AgensRoutes(route fiber.Router) {
	route.Get("/persons", controllers.GetPersons)
	route.Get("/person/:relation/:personName", controllers.GetPersonRelation)
	route.Post("/person", controllers.CreatePerson)
	route.Post("/connectPersons", controllers.ConnectPerson)
	route.Patch("/person/:personName", controllers.UpdatePerson)
	route.Delete("/person/:personName", controllers.DeletePerson)
	route.Delete("/persons", controllers.DeletePersons)
}
