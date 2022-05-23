package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pingcap-inc/tidb-example-golang/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Player struct {
	ID    string `gorm:"primaryKey;type:VARCHAR(36);column:id"`
	Coins int    `gorm:"column:coins"`
	Goods int    `gorm:"column:goods"`
}

func (*Player) TableName() string {
	return "player"
}

var db *gorm.DB

func dbInit() {
	// Configure the example database connection.
	dsn := "root:@tcp(127.0.0.1:4000)/test?charset=utf8mb4"
	initDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// AutoMigrate for player table
	initDB.AutoMigrate(&Player{})

	db = initDB
}

// createPlayers Create a player
func createPlayers(players []Player) error {
	for i := range players {
		players[i].ID = uuid.New().String()
	}
	return db.CreateInBatches(players, 100).Error
}

// getPlayerByID Get a player
func getPlayerByID(id string) (Player, error) {
	var testPlayer Player
	err := db.Find(&testPlayer, "id = ?", id).Error
	return testPlayer, err
}

// getPlayerByID Get a player
func getPlayerByLimit(limitSize int) ([]Player, error) {
	players := make([]Player, limitSize, limitSize)
	err := db.Limit(limitSize).Find(&players).Error
	return players, err
}

// getCount Count players amount.
func getCount() (int64, error) {
	playersCount := int64(0)
	err := db.Model(&Player{}).Count(&playersCount).Error
	return playersCount, err
}

func buyGoods(sellID, buyID string, amount, price int) error {
	return util.TiDBGormBegin(db, true, func(tx *gorm.DB) error {
		var sellPlayer, buyPlayer Player
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Find(&sellPlayer, "id = ?", sellID).Error; err != nil {
			return err
		}

		if sellPlayer.ID != sellID || sellPlayer.Goods < amount {
			return fmt.Errorf("sell player %s goods not enough", sellID)
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Find(&buyPlayer, "id = ?", buyID).Error; err != nil {
			return err
		}

		if buyPlayer.ID != buyID || buyPlayer.Coins < price {
			return fmt.Errorf("buy player %s coins not enough", buyID)
		}

		updateSQL := "UPDATE player set goods = goods + ?, coins = coins + ? WHERE id = ?"
		if err := tx.Exec(updateSQL, -amount, price, sellID).Error; err != nil {
			return err
		}

		if err := tx.Exec(updateSQL, amount, -price, buyID).Error; err != nil {
			return err
		}

		fmt.Println("\n[buyGoods]:\n    'trade success'")
		return nil
	})
}
