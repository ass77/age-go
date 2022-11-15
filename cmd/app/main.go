/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ass77/age-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupRoutes(app *fiber.App) {

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the root endpoint 😉",
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

	// Do cypher query to AGE with database/sql Tx API transaction conrol
	// fmt.Println("# Do cypher query with SQL API")
	// helpers.DoWithSqlAPI(dsn, graphName)

	// Do cypher query to AGE with Age API
	// fmt.Println("# Do cypher query with Age API")
	// helpers.DoWithAgeWrapper(dsn, graphName)

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
