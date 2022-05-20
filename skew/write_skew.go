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
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	openDB("mysql", "root:@tcp(127.0.0.1:4000)/test", func(db *sql.DB) {
		writeSkew(db)
	})
}

func openDB(driverName, dataSourceName string, runnable func(db *sql.DB)) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	runnable(db)
}

func writeSkew(db *sql.DB) {
	err := prepareData(db)
	if err != nil {
		panic(err)
	}

	waitingChan, waitGroup := make(chan bool), sync.WaitGroup{}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		err = askForLeave(db, waitingChan, 1, 1)
		if err != nil {
			panic(err)
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		err = askForLeave(db, waitingChan, 2, 2)
		if err != nil {
			panic(err)
		}
	}()

	waitGroup.Wait()
}

func askForLeave(db *sql.DB, waitingChan chan bool, goroutineID, doctorID int) error {
	txnComment := fmt.Sprintf("/* txn %d */ ", goroutineID)
	if goroutineID != 1 {
		txnComment = "\t" + txnComment
	}

	txn, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	fmt.Println(txnComment + "start txn")

	// Txn 1 should be waiting for txn 2 done
	if goroutineID == 1 {
		<-waitingChan
	}

	txnFunc := func() error {
		queryCurrentOnCall := "SELECT COUNT(*) AS `count` FROM `doctors` WHERE `on_call` = ? AND `shift_id` = ? FOR UPDATE"
		rows, err := txn.Query(queryCurrentOnCall, true, 123)
		if err != nil {
			return err
		}
		defer rows.Close()
		fmt.Println(txnComment + queryCurrentOnCall + " successful")

		count := 0
		if rows.Next() {
			err = rows.Scan(&count)
			if err != nil {
				return err
			}
		}
		rows.Close()

		if count < 2 {
			return fmt.Errorf("at least one doctor is on call")
		}

		shift := "UPDATE `doctors` SET `on_call` = ? WHERE `id` = ? AND `shift_id` = ?"
		_, err = txn.Exec(shift, false, doctorID, 123)
		if err == nil {
			fmt.Println(txnComment + shift + " successful")
		}
		return err
	}

	err = txnFunc()
	if err == nil {
		txn.Commit()
		fmt.Println("[runTxn] commit success")
	} else {
		txn.Rollback()
		fmt.Printf("[runTxn] got an error, rollback: %+v\n", err)
	}

	// Txn 2 done, let txn 1 run again
	if goroutineID == 2 {
		waitingChan <- true
	}

	return nil
}

func prepareData(db *sql.DB) error {
	err := createDoctorTable(db)
	if err != nil {
		return err
	}

	err = createDoctor(db, 1, "Alice", true, 123)
	if err != nil {
		return err
	}
	err = createDoctor(db, 2, "Bob", true, 123)
	if err != nil {
		return err
	}
	err = createDoctor(db, 3, "Carol", false, 123)
	if err != nil {
		return err
	}
	return nil
}

func createDoctorTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `doctors` (" +
		"    `id` int(11) NOT NULL," +
		"    `name` varchar(255) DEFAULT NULL," +
		"    `on_call` tinyint(1) DEFAULT NULL," +
		"    `shift_id` int(11) DEFAULT NULL," +
		"    PRIMARY KEY (`id`)," +
		"    KEY `idx_shift_id` (`shift_id`)" +
		"  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin")
	return err
}

func createDoctor(db *sql.DB, id int, name string, onCall bool, shiftID int) error {
	_, err := db.Exec("INSERT INTO `doctors` (`id`, `name`, `on_call`, `shift_id`) VALUES (?, ?, ?, ?)",
		id, name, onCall, shiftID)
	return err
}
