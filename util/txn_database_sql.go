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
	"context"
	"database/sql"
)

type TiDBSqlTx struct {
	*sql.Tx
	conn        *sql.Conn
	pessimistic bool
}

func TiDBSqlBegin(db *sql.DB, pessimistic bool) (*TiDBSqlTx, error) {
	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	if pessimistic {
		_, err = conn.ExecContext(ctx, "set @@tidb_txn_mode=?", "pessimistic")
	} else {
		_, err = conn.ExecContext(ctx, "set @@tidb_txn_mode=?", "optimistic")
	}
	if err != nil {
		return nil, err
	}
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &TiDBSqlTx{
		conn:        conn,
		Tx:          tx,
		pessimistic: pessimistic,
	}, nil
}

func (tx *TiDBSqlTx) Commit() error {
	defer tx.conn.Close()
	return tx.Tx.Commit()
}

func (tx *TiDBSqlTx) Rollback() error {
	defer tx.conn.Close()
	return tx.Tx.Rollback()
}
