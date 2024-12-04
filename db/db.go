package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {

	connStr := "postgresql://postgres.wkavzgwalsnnfvmqbxbf:akusukangoding@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Failed to parse db config %v", err)
	}

	config.ConnConfig.ConnectTimeout = 5 * time.Second
	config.MaxConnLifetime = 2 * time.Minute
	config.MaxConnIdleTime = 1 * time.Minute

	config.ConnConfig.StatementCacheCapacity = 100

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to create db pool %v", err)
	}

	Pool.Config().ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("DB failed ping: %v", err)
	}

	fmt.Println("DB connected")
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
		fmt.Println("Database connection closed")
	}
}
