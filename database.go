package main

import (
	"context"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
)

var ClickHouseConn *ch.DB

func createConnection() {
	ctx := context.Background()
	ClickHouseConn := ch.Connect(
		ch.WithDatabase("test"),
		ch.WithAutoCreateDatabase(true),
	)
	
		// db := ch.Connect(
		// 	ch.WithDatabase(dbName),
		// 	ch.WithAutoCreateDatabase(true),
		// )
	
	ClickHouseConn.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))
	
	if err := ClickHouseConn.Ping(ctx); err != nil {
		panic(err)
	}
	
}