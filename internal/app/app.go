package app

import (
	"context"

	"github.com/Thanhbinh1905/go-training-bank/internal/api"
	"github.com/Thanhbinh1905/go-training-bank/internal/config"
	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/service"
	"github.com/Thanhbinh1905/go-training-bank/pkg/connect"
	"github.com/Thanhbinh1905/go-training-bank/pkg/logger"
)

func Run(cfg *config.Config) {
	log := logger.InitLogger("go-training-bank", "logs/.log")
	defer log.Sync()

	pool, err := connect.Postgres(context.Background(), cfg.DatabaseURL, log)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	store := db.NewStore(pool)
	service := service.NewService(store)
	server := api.NewServer(service)

	if err := server.Start(":8080"); err != nil {
		panic(err)
	}
}
