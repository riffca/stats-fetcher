package main

import (
	"context"
	//"fmt"
	"time"

	"github.com/uptrace/go-clickhouse/ch"
	//"github.com/uptrace/go-clickhouse/chdebug"
)


type Lead struct {
	ch.CHModel `ch:"partition:toYYYYMM(time)"`

	ID   uint64
	Ip string    `ch:",lc"`
	Time time.Time `ch:",pk"`
}

func saveLeadStats(Ip string) {
	ctx := context.Background()
	src := &Lead{ID: 1, Ip: Ip, Time: time.Now()}
	if _, err := ClickHouseConn.NewInsert().Model(src).Exec(ctx); err != nil {
		panic(err)
	}
}

func getLeadStatsByIp(leadIp int) {
	dest := new(Lead)
	if err := ClickHouseConn.NewSelect().Model(dest).Where("ip = ?", leadIp).Limit(1).Scan(ctx); err != nil {
		panic(err)
	}
}




// type Model struct {
// 	ch.CHModel `ch:"partition:toYYYYMM(time)"`

// 	ID   uint64
// 	Text string    `ch:",lc"`
// 	Time time.Time `ch:",pk"`
// }

// func oops() {
// 	ctx := context.Background()

// 	ClickHouseConn := ch.Connect(ch.WithDatabase("test"))
// 	ClickHouseConn.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))

// 	if err := ClickHouseConn.Ping(ctx); err != nil {
// 		panic(err)
// 	}

// 	var num int
// 	if err := ClickHouseConn.QueryRowContext(ctx, "SELECT 123").Scan(&num); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(num)

// 	if err := ClickHouseConn.ResetModel(ctx, (*Model)(nil)); err != nil {
// 		panic(err)
// 	}

// 	src := &Model{ID: 1, Text: "hello", Time: time.Now()}
// 	if _, err := ClickHouseConn.NewInsert().Model(src).Exec(ctx); err != nil {
// 		panic(err)
// 	}

// 	dest := new(Model)
// 	if err := ClickHouseConn.NewSelect().Model(dest).Where("id = ?", src.ID).Limit(1).Scan(ctx); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(dest)
// }