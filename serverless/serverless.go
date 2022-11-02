package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

const (
	tlsConfigKeyName = "tidb"
	host             = "{serverless tier host}"
	port             = 4000
	username         = "{serverless tier username}"
	password         = "{serverless tier password}"
)

func main() {
	mysql.RegisterTLSConfig(tlsConfigKeyName, &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: host,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/test?tls=%s",
		username, password, host, port, tlsConfigKeyName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT VERSION()")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		version := ""
		rows.Scan(&version)
		fmt.Println(version)
	}
}
