package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"ordering-server/config"
	"os"
	"testing"
)

var testQueries Store

func TestMain(m *testing.M) {
	configEnv, err := config.LoadConfig("../../..")
	if err != nil {
		fmt.Println("Cannot load configurations", "error", err)
		return
	}
	conn, err := pgxpool.New(context.Background(), configEnv.DBUrl)
	if err != nil {
		log.Fatal("Cannot connect to database", "error", err)
		return
	}

	testQueries = NewStore(conn)
	os.Exit(m.Run())
}
