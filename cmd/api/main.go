package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"ordering-server/config"
	"ordering-server/internal/database"
	"ordering-server/internal/server"
)

func main() {
	configEnv, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot load configurations", "error", err)
		return
	}
	fmt.Println("Configurations loaded")
	store, err := InitDb(configEnv.DBUrl)
	s, err := server.NewServer(store)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = s.RunServer(configEnv.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}

func InitDb(DBUrl string) (*database.Store, error) {
	conn, err := pgxpool.New(context.Background(), DBUrl)
	if err != nil {
		return nil, err
	}
	return database.NewStore(conn), nil
}
