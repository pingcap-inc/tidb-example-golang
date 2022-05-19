# batch

This is an example of a batch update or delete using TiDB and Golang.

## Running way

1. Makefile(recommend)

- Run `make all`.

- Batch delete: Run `make bulk-delete`
- Batch update: Run `make bulk-update`
- 
2. Manual

- Batch delete:
  - Enter `delete` catalog.
  - Run `tiup demo bookshop prepare --drop-tables` to create the data structure and data on TiDB.
  - Run `go build -o bin/batch-delete`.
  - Run `./bin/batch-delete`.

- Batch update:
    - Enter `update` catalog.
    - Run `tiup demo bookshop prepare --drop-tables` to create the data structure and data on TiDB.
    - Run `mysql --host 127.0.0.1 --port 4000 -u root<add_attr_ten_point.sql` to add `ten_point` field.
    - Run `go build -o bin/batch-update`.
    - Run `./bin/batch-update`.

## Code

- [Batch Update Example](./update/update.go)
- [Batch Delete Example](./delete/delete.go)
