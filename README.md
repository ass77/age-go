# age AGType parser and driver support for Golang

AGType parser and driver support for [Apache AGE](https://age.apache.org/), graph extention for PostgreSQL.

### Features

- Unmarshal AGE result data(AGType) to Vertex, Edge, Path
- Cypher query support for 3rd. Party sql driver (enables to use cypher queries directly)

### Prerequisites

- over Go 1.16
- This module runs on golang standard api [database/sql](https://golang.org/pkg/database/sql/) and [antlr4-python3](https://github.com/antlr/antlr4/tree/master/runtime/Go/antlr)

### Go get

```
go get github.com/apache/age/drivers/golang

```

Check [latest version](https://github.com/apache/age/releases)

### For more information about [Apache AGE](https://age.apache.org/)

- Apache Age : https://age.apache.org/
- Github : https://github.com/apache/age
- Document : https://age.apache.org/docs/

### Check AGE loaded on your PostgreSQL

Connect to your containerized Postgres instance and then run the following commands:

```(sql)
# psql
CREATE EXTENSION age;
LOAD 'age';
SET search_path = ag_catalog, "$user", public;
```

### Test

```
cd age
go test . -v

```

### Application

- Usage 1: using database/sql API and Cypher execution function 'ExecCypher'
  Sample : [cmd/app/sql_api_sample.go](cmd/app/sql_api_sample.go)

- Usage 2: using Age Wrapper
  Sample : [cmd/app/age_wrapper_sample.go](cmd/app/age_wrapper_sample.go)

- Run Samples : [cmd/app/main.go](cmd/app/main.go)

### Setting up Postgres11 with AGE extension

Check [this issue](https://github.com/apache/age/issues/347)
