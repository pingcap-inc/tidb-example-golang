# tidb-example-golang

## Outline

It's an example for Golang and TiDB. Contains subproject:

- [gorm example](#gorm)
- [go-sql-driver/mysql example](#plain-java-hibernate)

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
