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
	"flag"
	"sync"
)

func main() {
	optimistic, alice, bob := parseParams()

	openDB("mysql", "root:@tcp(127.0.0.1:4000)/bookshop?charset=utf8mb4", func(db *sql.DB) {
		prepareData(db, optimistic)
		buy(db, optimistic, alice, bob)
	})
}

func buy(db *sql.DB, optimistic bool, alice, bob int) {
	buyFunc := buyOptimistic
	if !optimistic {
		buyFunc = buyPessimistic
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		buyFunc(db, 1, 1000, 1, 1, bob)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		buyFunc(db, 2, 1001, 1, 2, alice)
	}()

	wg.Wait()
}

func openDB(driverName, dataSourceName string, runnable func(db *sql.DB)) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	runnable(db)
}

func parseParams() (optimistic bool, alice, bob int) {
	flag.BoolVar(&optimistic, "o", false, "transaction is optimistic")
	flag.IntVar(&alice, "a", 4, "Alice bought num")
	flag.IntVar(&bob, "b", 6, "Bob bought num")

	flag.Parse()

	return optimistic, alice, bob
}
