# txn

This is an example of the optimistic and pessimistic transaction using TiDB and Golang.

## Running way

1. Makefile(recommend)

- Pessimistic transaction: Run `make pessimistic`
- Optimistic transaction: Run `make optimistic`

2. Manual

- Pessimistic transaction
  - Run `go build -o bin/txn` to build binary file.
  - Run `tiup demo bookshop prepare --drop-tables --books 0 --authors 0 --orders 0 --ratings 0 --users 0` to create the data structure only.
  - Run `./bin/txn -a 4 -b 6` to check not oversell example output.
  - Run `tiup demo bookshop prepare --drop-tables --books 0 --authors 0 --orders 0 --ratings 0 --users 0` to create the data structure again.
  - Run `./bin/txn -a 4 -b 7` to check oversell example output.

- Optimistic transaction
    - Run `go build -o bin/txn` to build binary file.
    - Run `tiup demo bookshop prepare --drop-tables --books 0 --authors 0 --orders 0 --ratings 0 --users 0` to create the data structure only.
    - Run `./bin/txn -o -a 4 -b 6` to check not oversell example output.
    - Run `tiup demo bookshop prepare --drop-tables --books 0 --authors 0 --orders 0 --ratings 0 --users 0` to create the data structure again.
    - Run `./bin/txn -o -a 4 -b 7` to check oversell example output.

## Code

- [Main Entry](./txn.go)
- [Transaction Helper](./helper.go)