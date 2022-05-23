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
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

type TiDBVar struct {
	Name  string `gorm:"column:Variable_name"`
	Value string `gorm:"column:Value"`
}

func TestTiDBGormBegin(t *testing.T) {
	testDB := createTestGormDB()

	err := TiDBGormBegin(testDB, true, func(tx *gorm.DB) error {
		txnVar := TiDBVar{}
		result := tx.Raw("SHOW SESSION VARIABLES LIKE 'tidb_txn_mode'").Scan(&txnVar)
		if result.Error != nil {
			t.Error(result.Error)
		}

		if txnVar.Value != "pessimistic" {
			t.Error(fmt.Errorf("showld be pessimistic"))
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	err = TiDBGormBegin(testDB, false, func(tx *gorm.DB) error {
		txnVar := TiDBVar{}
		result := tx.Raw("SHOW SESSION VARIABLES LIKE 'tidb_txn_mode'").Scan(&txnVar)
		if result.Error != nil {
			t.Error(result.Error)
		}

		if txnVar.Value != "optimistic" {
			t.Error(fmt.Errorf("showld be optimistic"))
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func createTestGormDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:4000)/test?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}
