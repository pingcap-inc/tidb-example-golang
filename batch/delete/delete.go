// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:4000)/bookshop")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	affectedRows := int64(-1)
	startTime := time.Date(2022, 04, 15, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2022, 04, 15, 0, 15, 0, 0, time.UTC)

	for affectedRows != 0 {
		affectedRows, err = deleteBatch(db, startTime, endTime)
		if err != nil {
			panic(err)
		}
	}
}

// deleteBatch delete at most 1000 lines per batch
func deleteBatch(db *sql.DB, startTime, endTime time.Time) (int64, error) {
	bulkUpdateSql := fmt.Sprintf("DELETE FROM `bookshop`.`ratings` WHERE `rated_at` >= ? AND  `rated_at` <= ? LIMIT 1000")
	result, err := db.Exec(bulkUpdateSql, startTime, endTime)
	if err != nil {
		return -1, err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	fmt.Printf("delete %d data\n", affectedRows)
	return affectedRows, nil
}
