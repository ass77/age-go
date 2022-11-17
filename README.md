<div align="center">
  <a href="https://tgcollective.xyz">
    <img src="https://www.tgcollective.xyz/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Frexxie-banner.227d942b.webp&w=3840&q=75" alt="tgc_fam" width="800" height="200">
  </a>
  <h3 align="center">Graph DB with postgreSQL AGE extension server with GoFiber</h3>
  <h3 align="center">Created by: <a href="https://github.com/ass77">ass77</a></h3>
</div>

## Features

- RESTful endpoints (CRUD) vertexes, edges, paths
- Unmarshal AGE result data(AGType) to Vertex, Edge, Path
- Cypher query support for 3rd. Party sql driver (enables to use cypher queries directly)

## Prerequisites

- over Go 1.16
- make
- Postgres 11.18
- This module runs on golang standard api [database/sql](https://golang.org/pkg/database/sql/) and [antlr4-python3](https://github.com/antlr/antlr4/tree/master/runtime/Go/antlr)

## How To

### Setting up Postgres11 with AGE extension

- Check this [issue](https://github.com/apache/age/issues/347) out!

### Install dependencies

```
make tidy
```

### Test Run

```
make test
```

### Run server

```
export APP_ENV=local
make run
```

### Endpoints

- import [insomnia docs](./insomnia/Insomnia_2022-11-17.json) to the your insomnia.

## Extension vs No Extension

- [AGE extension](./age.sql) only supports [Postgres 11](https://github.com/apache#:~:text=Apache%20AGE%20is%20currently%20being%20developed%20for%20the%20PostgreSQL%2012%20release%20and%20will%20support%20PostgreSQL%2013%20and%20all%20the%20future%20releases%20of%20PostgreSQL.) at the moment.
- [With Recursive Usage](./withRecursive.sql) is supported by any postgres version.

## For more information about [Apache AGE](https://age.apache.org/)

- Apache Age : https://age.apache.org/
- Github : https://github.com/apache/age
- Document : https://age.apache.org/docs/
- Check [latest version](https://github.com/apache/age/releases)
