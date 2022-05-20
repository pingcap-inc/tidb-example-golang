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
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:4000)/bookshop")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookID, userID := updateBatch(db, true, 0, 0)
	fmt.Println("first time batch update success")
	for {
		time.Sleep(time.Second)
		bookID, userID = updateBatch(db, false, bookID, userID)
		fmt.Printf("batch update success, [bookID] %d, [userID] %d\n", bookID, userID)
	}
}

// updateBatch select at most 1000 lines data to update score
func updateBatch(db *sql.DB, firstTime bool, lastBookID, lastUserID int64) (bookID, userID int64) {
	// select at most 1000 primary keys in five-point scale data
	var err error
	var rows *sql.Rows

	if firstTime {
		rows, err = db.Query("SELECT `book_id`, `user_id` FROM `bookshop`.`ratings` " +
			"WHERE `ten_point` != true ORDER BY `book_id`, `user_id` LIMIT 1000")
	} else {
		rows, err = db.Query("SELECT `book_id`, `user_id` FROM `bookshop`.`ratings` "+
			"WHERE `ten_point` != true AND `book_id` > ? AND `user_id` > ? "+
			"ORDER BY `book_id`, `user_id` LIMIT 1000", lastBookID, lastUserID)
	}

	if err != nil || rows == nil {
		panic(fmt.Errorf("error occurred or rows nil: %+v", err))
	}

	// joint all id with a list
	var idList []interface{}
	for rows.Next() {
		var tempBookID, tempUserID int64
		if err := rows.Scan(&tempBookID, &tempUserID); err != nil {
			panic(err)
		}
		idList = append(idList, tempBookID, tempUserID)
		bookID, userID = tempBookID, tempUserID
	}

	bulkUpdateSql := fmt.Sprintf("UPDATE `bookshop`.`ratings` SET `ten_point` = true, "+
		"`score` = `score` * 2 WHERE (`book_id`, `user_id`) IN (%s)", placeHolder(len(idList)))
	db.Exec(bulkUpdateSql, idList...)

	return bookID, userID
}

// placeHolder format SQL place holder
func placeHolder(n int) string {
	holderList := make([]string, n/2, n/2)
	for i := range holderList {
		holderList[i] = "(?,?)"
	}
	return strings.Join(holderList, ",")
}
