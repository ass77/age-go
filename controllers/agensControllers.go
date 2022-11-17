package controllers

import (
	"log"

	"github.com/ass77/age-go/age"
	"github.com/ass77/age-go/config"
	"github.com/ass77/age-go/models"
	"github.com/gofiber/fiber/v2"
)

var db = config.ConnectDB()
var graphName string = "working_person"

func GetPersons(c *fiber.Ctx) error {

	relation := c.Params("relation")

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	cursor, err := age.ExecCypher(tx, graphName, 3, "MATCH (a:Person)-[l:%s]-(b:Person) RETURN a, l, b", relation)

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

		allData = append(allData, models.Vertex{
			V1:   v1.String(),
			Edge: edge.String(),
			V2:   v2.String(),
		})
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"vertex": allData,
	})

}

func GetPersonRelation(c *fiber.Ctx) error {

	// Get id from url params
	name := c.Params("personName")
	relation := c.Params("relation")

	// Confirm graph_path created

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// find vertices with Cypher based on name
	cursor, err := age.ExecCypher(tx, graphName, 3, "MATCH (a:Person)-[l:%s]-(b:Person) WHERE a.name = '%s' RETURN a, l, b", relation, name)

	var data []models.Vertex

	for cursor.Next() {
		row, err := cursor.GetRow()
		if err != nil {
			panic(err)
		}
		v1 := row[0].(*age.Vertex)
		edge := row[1].(*age.Edge)
		v2 := row[2].(*age.Vertex)

		data = append(data, models.Vertex{
			V1:   v1.String(),
			Edge: edge.String(),
			V2:   v2.String(),
		})
	}

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"vertex": data,
	})

}

func CreatePersonNode(c *fiber.Ctx) error {

	// Confirm graph_path created

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

	for vertex.Next() {
		row, err := vertex.GetRow()
		if err != nil {
			panic(err)
		}
		v := row[0].(*age.Vertex)
		log.Println("VERTEX ", v)
	}

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

func ConnectPersonNode(c *fiber.Ctx) error {

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	var connection models.ConnectPerson
	var data []models.Vertex
	err = c.BodyParser(&connection)
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 3, "MATCH (a:Person {name: '%s'}), (b:Person {name: '%s'}) CREATE (a)-[r:%s {weight: %f}]->(b) RETURN a, r, b", connection.PersonA, connection.PersonB, connection.Relation, connection.Weight)

	for vertex.Next() {
		row, err := vertex.GetRow()
		if err != nil {
			panic(err)
		}
		v1 := row[0].(*age.Vertex)
		edge := row[1].(*age.Edge)
		v2 := row[2].(*age.Vertex)

		data = append(data, models.Vertex{
			V1:   v1.String(),
			Edge: edge.String(),
			V2:   v2.String(),
		})
	}

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"vertex": data,
	})
}

func UpdatePersonNode(c *fiber.Ctx) error {

	name := c.Params("personName")

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
	vertex, err := age.ExecCypher(tx, graphName, 0, "MATCH (n:Person {name: '%s'}) SET n.weight = %f n.role = '%s' RETURN *", name, person.Weight, person.Role)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Person updated",
		"vertex":  vertex,
	})

}

func DeletePersonNode(c *fiber.Ctx) error {

	name := c.Params("personName")

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 1, "MATCH (n:Person {name: '%s'}) DETACH DELETE n RETURN n", name)

	for vertex.Next() {
		row, err := vertex.GetRow()
		if err != nil {
			panic(err)
		}
		v := row[0].(*age.Vertex)
		log.Println("VERTEX ", v)
	}

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Person deleted",
		"vertex":  vertex,
	})

}

func DeletePersonsNodes(c *fiber.Ctx) error {

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	vertex, err := age.ExecCypher(tx, graphName, 1, "MATCH (n:Person) DETACH DELETE n RETURN n")

	for vertex.Next() {
		row, err := vertex.GetRow()

		log.Println("ROW ", row)

		if err != nil {
			panic(err)
		}

		v := row[0].(*age.Vertex)
		log.Println("VERTEX ", v)
	}

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "All Persons deleted",
		"vertex":  vertex,
	})
}
