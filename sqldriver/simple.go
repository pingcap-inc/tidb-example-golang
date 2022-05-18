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

func getPlayer(db *sql.DB, id string) (*Player, error) {
	player := Player{}

	rows, err := db.Query(GetPlayerSQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		scanErr := rows.Scan(&player.ID, &player.Coins, &player.Goods)
		if scanErr != nil {
			return nil, scanErr
		}

		return &player, nil
	}

	return nil, fmt.Errorf("can not found player")
}

func bulkInsertRandomPlayers(db *sql.DB, total, batchSize int) error {
	db.Begin()
}
