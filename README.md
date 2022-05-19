# tidb-example-golang

## Outline

It's an example for Golang and TiDB. Contains subproject:

- [gorm example](#gorm-example)
- [go-sql-driver/mysql example](#go-sql-drivermysql-example)
- [http example](#http-example)

We use different frameworks or libraries to implement similar processes to reflect
the solution for connecting to TiDB in different environments

This is a process about the game, each player has two attributes,
`coins` and `goods`, and each player has their own unique `id` as an identifier.
Players can trade with each other, provided that the `coins` and `goods` are sufficient

The process is as follows:

1. Create a player
2. Create some players
3. Read players amount
4. Read some players attributes
5. Two player trade with insufficient coins or goods
6. Two player trade with sufficient coins or goods

## Dependency

- [Golang SDK](https://go.dev/)
- mysql client

## gorm example

It's an example used [gorm](https://gorm.io/docs/index.html) to connect TiDB.

### Running way

1. Makefile(recommend)
    - Run `make gorm-example`

2. Manual
    - Into `gorm`
    - Run `go build -o bin/gorm-example`
    - Run `./bin/gorm-example`

### Expected output

gorm-example [expected output](./Expected-Output.md#gorm)

### Code

- [Demo](./gorm/gorm.go)

## go-sql-driver/mysql example

It's an example used [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) to connect TiDB.

### Running way

1. Makefile(recommend)
    - Run `make sql-driver-example`

2. Manual
    - Run [create table sql](./sqldriver/sql/dbinit.sql) in your TiDB
    - Into `sqldriver`
    - Run `go build -o bin/sql-driver-example`
    - Run `./bin/sql-driver-example`

### Expected output

go-sql-driver/mysql [expected output](./Expected-Output.md#sqldriver)

### Code

- [Initial SQL](./sqldriver/sql/dbinit.sql)
- [Data Access Functions](./sqldriver/dao.go)
- [SQL Strings](./sqldriver/sql.go)
- [Main Entry](./sqldriver/sqldriver.go)

## http example

It's an example service used gorm to connect TiDB.
Provide a group of HTTP Restful interface.

### Running way

1. Makefile(recommend)
   1. First terminal
       - Run `make http-example`
   2. Second terminal
       - Run `make request`
   3. You can quit first terminal to stop service

2. Manual
    - Into `http`
    - Run `go build -o bin/http`
    - Run `./bin/http`
    - Request:
      - Option 1 (recommend):
        - Input [Request Collection](./http/Player.postman_collection.json) to [Postman](https://www.postman.com/)
        - Request by Postman application
      - Option 2:
        - Using [script](./http/request.sh) to request. It's based on `curl`
      - Option 3:
        - Request HTTP Restful interface by other way

### Expected output

1. request [expected output](./Expected-Output.md#http-request)
2. service [expected output](./Expected-Output.md#http-service)

### Code

- [Main Entry](./http/http.go)
- [Logic](./http/service.go)

## More Demo

- [Transaction](./txn/README.md)
- [Batch Delete/Update](./batch/README.md)
- [Write Skew](./skew/README.md)
