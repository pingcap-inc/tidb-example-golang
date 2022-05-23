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

package util

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestTiDBSqlBegin(t *testing.T) {
	testDB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:4000)/test?charset=utf8mb4")
	if err != nil {
		t.Error(err)
	}

	tidbSQLTxnTest(testDB, true, true, t)
	tidbSQLTxnTest(testDB, true, false, t)
	tidbSQLTxnTest(testDB, false, true, t)
	tidbSQLTxnTest(testDB, false, false, t)
}

func tidbSQLTxnTest(testDB *sql.DB, pessimistic, rollback bool, t *testing.T) {
	txn, err := TiDBSqlBegin(testDB, pessimistic)
	if err != nil {
		t.Error(err)
	}

	result, err := txn.Query("SHOW SESSION VARIABLES LIKE 'tidb_txn_mode'")
	if err != nil {
		t.Error(err)
	}

	for result.Next() {
		name, value := "", ""
		err = result.Scan(&name, &value)
		if err != nil {
			t.Error(err)
		}

		if name == "tidb_txn_mode" {
			if pessimistic && value != "pessimistic" {
				t.Error(fmt.Errorf("showld be pessimistic"))
			} else if !pessimistic && value != "optimistic" {
				t.Error(fmt.Errorf("showld be optimistic"))
			}
		}
	}

	if rollback {
		err = txn.Rollback()
	} else {
		err = txn.Commit()
	}

	if err != nil {
		t.Error(err)
	}
}
