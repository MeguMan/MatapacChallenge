package main

import (
	"context"
	"fmt"
	"github.com/MeguMan/MatapacChallenge/internal/services/chainstask_service"
	"github.com/MeguMan/MatapacChallenge/internal/services/telegram_service"
	"github.com/MeguMan/MatapacChallenge/internal/storage"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()

	chainstackService := chainstask_service.New(os.Getenv("CHAINSTACK_URL"))
	storageRepo, err := storage.New(db)
	if err != nil {
		log.Fatalf("Failed to create storage: %v", err)
	}

	telegramService, err := telegram_service.New(ctx, chainstackService, storageRepo, os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to create telegram service: %v", err)
	}

	if err = telegramService.Handle(ctx); err != nil {
		log.Fatal(err)
	}
}
