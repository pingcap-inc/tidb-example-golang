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
    "math/rand"
    "strings"

    "github.com/google/uuid"
    "github.com/pingcap-inc/tidb-example-golang/util"
)

type Player struct {
    ID    string
    Coins int
    Goods int
}

// createPlayer create a player
func createPlayer(db *sql.DB, player Player) error {
    _, err := db.Exec(CreatePlayerSQL, player.ID, player.Coins, player.Goods)
    return err
}

// getPlayer get a player
func getPlayer(db *sql.DB, id string) (Player, error) {
    var player Player

    rows, err := db.Query(GetPlayerSQL, id)
    if err != nil {
        return player, err
    }
    defer rows.Close()

    if rows.Next() {
        err = rows.Scan(&player.ID, &player.Coins, &player.Goods)
        if err == nil {
            return player, nil
        } else {
            return player, err
        }
    }

    return player, fmt.Errorf("can not found player")
}

// getPlayerByLimit get players by limit
func getPlayerByLimit(db *sql.DB, limit int) ([]Player, error) {
    var players []Player

    rows, err := db.Query(GetPlayerByLimitSQL, limit)
    if err != nil {
        return players, err
    }
    defer rows.Close()

    for rows.Next() {
        player := Player{}
        err = rows.Scan(&player.ID, &player.Coins, &player.Goods)
        if err == nil {
            players = append(players, player)
        } else {
            return players, err
        }
    }

    return players, nil
}

// bulk-insert players
func bulkInsertPlayers(db *sql.DB, players []Player, batchSize int) error {
    tx, err := util.TiDBSqlBegin(db, true)
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare(buildBulkInsertSQL(batchSize))
    if err != nil {
        return err
    }

    defer stmt.Close()

    for len(players) > batchSize {
        if _, err := stmt.Exec(playerToArgs(players[:batchSize])...); err != nil {
            tx.Rollback()
            return err
        }

        players = players[batchSize:]
    }

    if len(players) != 0 {
        if _, err := tx.Exec(buildBulkInsertSQL(len(players)), playerToArgs(players)...); err != nil {
            tx.Rollback()
            return err
        }
    }

    if err := tx.Commit(); err != nil {
        tx.Rollback()
        return err
    }

    return nil
}

func getCount(db *sql.DB) (int, error) {
    count := 0

    rows, err := db.Query(GetCountSQL)
    if err != nil {
        return count, err
    }

    defer rows.Close()

    if rows.Next() {
        if err := rows.Scan(&count); err != nil {
            return count, err
        }
    }

    return count, nil
}

func tradeCheck(tx *util.TiDBSqlTx, sellID, buyID string, amount, price int) error {
    var sellPlayer, buyPlayer Player
    stmt, err := tx.Prepare(GetPlayerWithLockSQL)
    if err != nil {
        return err
    }
    defer stmt.Close()

    sellRows, err := stmt.Query(sellID)
    if err != nil {
        return err
    }
    defer sellRows.Close()

    if sellRows.Next() {
        if err := sellRows.Scan(&sellPlayer.ID, &sellPlayer.Coins, &sellPlayer.Goods); err != nil {
            return err
        }
    }
    sellRows.Close()

    if sellPlayer.ID != sellID || sellPlayer.Goods < amount {
        return fmt.Errorf("sell player %s goods not enough", sellID)
    }

    buyRows, err := stmt.Query(buyID)
    if err != nil {
        return err
    }
    defer buyRows.Close()

    if buyRows.Next() {
        if err := buyRows.Scan(&buyPlayer.ID, &buyPlayer.Coins, &buyPlayer.Goods); err != nil {
            return err
        }
    }
    buyRows.Close()

    if buyPlayer.ID != buyID || buyPlayer.Coins < price {
        return fmt.Errorf("buy player %s coins not enough", buyID)
    }

    return nil
}

func tradeUpdate(tx *util.TiDBSqlTx, sellID, buyID string, amount, price int) error {
    updateStmt, err := tx.Prepare(UpdatePlayerSQL)
    if err != nil {
        return err
    }
    defer updateStmt.Close()

    if _, err := updateStmt.Exec(-amount, price, sellID); err != nil {
        return err
    }

    if _, err := updateStmt.Exec(amount, -price, buyID); err != nil {
        return err
    }

    return nil
}

func buyGoods(db *sql.DB, sellID, buyID string, amount, price int) error {
    tx, err := util.TiDBSqlBegin(db, true)
    if err != nil {
        return err
    }

    buyExec := func() error {
        if err := tradeCheck(tx, sellID, buyID, amount, price); err != nil {
            return err
        }

        return tradeUpdate(tx, sellID, buyID, amount, price)
    }

    err = buyExec()
    if err == nil {
        fmt.Println("\n[buyGoods]:\n    'trade success'")
        tx.Commit()
    } else {
        tx.Rollback()
    }

    return err
}

func playerToArgs(players []Player) []interface{} {
    var args []interface{}
    for _, player := range players {
        args = append(args, player.ID, player.Coins, player.Goods)
    }
    return args
}

func buildBulkInsertSQL(amount int) string {
    return CreatePlayerSQL + strings.Repeat(",(?,?,?)", amount-1)
}

func randomPlayers(amount int) []Player {
    players := make([]Player, amount, amount)
    for i := 0; i < amount; i++ {
        players[i] = Player{
            ID:    uuid.New().String(),
            Coins: rand.Intn(10000),
            Goods: rand.Intn(10000),
        }
    }

    return players
}
