package main

import (
	"context"
	"math/rand"
	"os"
	"team-maker/api"
	"team-maker/utils"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	err := godotenv.Load()
	if err != nil {
		zap.L().Fatal("Failed to get env var", zap.Error(err))
	}
	err = utils.S3Init()
	if err != nil {
		zap.L().Fatal("Failed to init S3 client", zap.Error(err))
	}

	var db *pgxpool.Pool

	dbUrl := os.Getenv("DATABASE_URL")

	ctx := context.Background()
	db, err = pgxpool.Connect(ctx, dbUrl)
	if err != nil {
		zap.L().Fatal("DB Connection", zap.Error(err))
	}

	defer func() {
		db.Close()
	}()

	s := api.MakeServer(db)
	s.RunServer()
}
