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
	"os"

	helpers "github.com/ass77/age-go/helpers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// var dsn string = "host={host} port={port} dbname={dbname} user={username} password={password} sslmode=disable"

// var graphName string = "{graph_path}"

// TODO generic graphDB -> use metadata to create vertexes and match edges
func main() {

	var ENV_TYPE = os.Getenv("APP_ENV")
	var dsn string

	if ENV_TYPE == "local" {
		// load the env file
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")

		}

		dsn = os.Getenv("LOCAL_DSN")

	} else if ENV_TYPE == "development" {
		dsn = os.Getenv("DEVELOPMENT_DSN")

	} else if ENV_TYPE == "staging" {
		fmt.Println("staging")
		dsn = os.Getenv("STAGING_DSN")

	} else {
		fmt.Println("production")
		dsn = os.Getenv("PRODUCTION_DSN")

	}

	fmt.Println("ENV_TYPE: ", ENV_TYPE)
	fmt.Println("dsn: ", dsn)
	var graphName string = "working_person"

	// Do cypher query to AGE with database/sql Tx API transaction conrol
	fmt.Println("# Do cypher query with SQL API")
	helpers.DoWithSqlAPI(dsn, graphName)

	// Do cypher query to AGE with Age API
	fmt.Println("# Do cypher query with Age API")
	helpers.DoWithAgeWrapper(dsn, graphName)

}
