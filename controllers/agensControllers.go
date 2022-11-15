package controllers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ass77/age-go/age"
	"github.com/ass77/age-go/models"
	"github.com/gofiber/fiber/v2"
)

func GetPersons(c *fiber.Ctx) error {

	var dsn = os.Getenv("DSN")

	var graphName string = "working_person"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	cursor, err := age.ExecCypher(tx, graphName, 3, "MATCH (a:Person)-[l:workWith]-(b:Person) RETURN a, l, b")
	if err != nil {
		panic(err)
	}

	// initialize empty array
	var allData []models.Vertex

	count := 0

	for cursor.Next() {
		row, err := cursor.GetRow()
		if err != nil {
			panic(err)
		}
		count++
		v1 := row[0].(*age.Vertex)
		edge := row[1].(*age.Edge)
		v2 := row[2].(*age.Vertex)
		fmt.Println("ROW ", count, ">>", "\n\t", v1, "\n\t", edge, "\n\t", v2)

		// change v1 to string
		v1Str := fmt.Sprintf("%v", v1)

		edgeStr := fmt.Sprintf("%v", edge)

		// change v2 to string
		v2Str := fmt.Sprintf("%v", v2)

		// append to array
		allData = append(allData, models.Vertex{
			V1:   v1Str,
			Edge: edgeStr,
			V2:   v2Str,
		})
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Persons found",
		"data":    allData,
	})

}

func GetPerson(c *fiber.Ctx) error {

	// Get id from url params
	id := c.Params("personId")

	var dsn = os.Getenv("DSN")

	var graphName string = "working_person"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// find vertices with Cypher based on id
	vertex, err := age.ExecCypher(tx, graphName, 1, "MATCH (n:Person {name: '%s'}) RETURN n", id)

	// if vertex == nil {
	// return c.Status(404).JSON(fiber.Map{
	// "message": "No Person found",
	// })
	// }

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Person found",
		"vertex":  vertex,
	})

}

func CreatePerson(c *fiber.Ctx) error {

	var dsn = os.Getenv("DSN")

	var graphName string = "working_person"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// get name and weight from request body with models.Person
	var person models.Person
	err = c.BodyParser(&person)
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 1, "CREATE (n:Person {name: '%s', role:'%s', weight:%f}) RETURN n", person.Name, person.Role, person.Weight)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Person created",
		"vertex":  vertex,
	})
}

func ConnectPersons(c *fiber.Ctx) error {

	var dsn = os.Getenv("DSN")

	var graphName string = "working_person"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// get name and weight from request body with models.Person
	var person models.ConnectPerson
	err = c.BodyParser(&person)
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 3, "MATCH (a:Person {name: '%s'}), (b:Person {name: '%s'}) CREATE (a)-[r:workWith {weight: %f}]->(b) RETURN a, r, b", person.PersonA, person.PersonB, person.Weight)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": person.PersonA + " and " + person.PersonB + " connected",
		"vertex":  vertex,
	})
}

func UpdatePerson(c *fiber.Ctx) error {

	var dsn = os.Getenv("DSN")

	var graphName string = "working_person"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	var person models.Person
	err = c.BodyParser(&person)
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 0, "MATCH (n:Person {name: '%s'}) SET n.weight = %f n.role = '%s' RETURN *", person.Name, person.Weight, person.Role)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Person updated",
		"vertex":  vertex,
	})

}

func DeletePerson(c *fiber.Ctx) error {

	var dsn = os.Getenv("DSN")

	var graphName string = "working_person"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	name := c.Params("personName")

	fmt.Println(name, "name...............")

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 1, "MATCH (n:Person {name: '%s'}) DETACH DELETE n RETURN n", name)

	// if vertex == nil {
	// return c.Status(404).JSON(fiber.Map{
	// "message": "No Person found",
	// })
	// }

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Person deleted",
		"vertex":  vertex,
	})

}

func DeletePersons(c *fiber.Ctx) error {
	var dsn = os.Getenv("DSN")

	var graphName string = "working_person"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 1, "MATCH (n:Person) DETACH DELETE n RETURN n")

	// if vertex == nil {
	// return c.Status(404).JSON(fiber.Map{
	// "message": "No Person found",
	// })
	// }

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "All Persons deleted",
		"vertex":  vertex,
	})
}
