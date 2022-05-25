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

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 1. Configure the example database connection.
	dsn := "root:@tcp(127.0.0.1:4000)/test?charset=utf8mb4"
	openDB("mysql", dsn, func(db *sql.DB) {
		// 2. Run some simple examples.
		simpleExample(db)

		// 3. Getting further.
		tradeExample(db)
	})
}

func simpleExample(db *sql.DB) {
	// Create a player, who has a coin and a goods.
	err := createPlayer(db, Player{ID: "test", Coins: 1, Goods: 1})
	if err != nil {
		panic(err)
	}

	// Get a player.
	testPlayer, err := getPlayer(db, "test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("getPlayer: %+v\n", testPlayer)

	// Create players with bulk inserts. Insert 1919 players totally, with 114 players per batch.

	err = bulkInsertPlayers(db, randomPlayers(1919), 114)
	if err != nil {
		panic(err)
	}

	// Count players amount.
	playersCount, err := getCount(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("countPlayers: %d\n", playersCount)

	// Print 3 players.
	threePlayers, err := getPlayerByLimit(db, 3)
	if err != nil {
		panic(err)
	}
	for index, player := range threePlayers {
		fmt.Printf("print %d player: %+v\n", index+1, player)
	}
}

func tradeExample(db *sql.DB) {
	// Player 1: id is "1", has only 100 coins.
	// Player 2: id is "2", has 114514 coins, and 20 goods.
	player1 := Player{ID: "1", Coins: 100}
	player2 := Player{ID: "2", Coins: 114514, Goods: 20}

	// Create two players "by hand", using the INSERT statement on the backend.
	if err := createPlayer(db, player1); err != nil {
		panic(err)
	}
	if err := createPlayer(db, player2); err != nil {
		panic(err)
	}

	// Player 1 wants to buy 10 goods from player 2.
	// It will cost 500 coins, but player 1 cannot afford it.
	fmt.Println("\nbuyGoods:\n    => this trade will fail")
	if err := buyGoods(db, player2.ID, player1.ID, 10, 500); err == nil {
		panic("there shouldn't be success")
	}

	// So player 1 has to reduce the incoming quantity to two.
	fmt.Println("\nbuyGoods:\n    => this trade will success")
	if err := buyGoods(db, player2.ID, player1.ID, 2, 100); err != nil {
		panic(err)
	}
}

func openDB(driverName, dataSourceName string, runnable func(db *sql.DB)) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	runnable(db)
}
