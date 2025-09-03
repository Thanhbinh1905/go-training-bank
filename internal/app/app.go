package app

import (
	"context"

	"github.com/Thanhbinh1905/go-training-bank/internal/api"
	"github.com/Thanhbinh1905/go-training-bank/internal/config"
	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/service"
	"github.com/Thanhbinh1905/go-training-bank/pkg/logger"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	log := logger.InitLogger("go-training-bank", "logs/.log")
	defer log.Sync()

	pool, err := db.Connect(context.Background(), cfg.DBSource, log)
	if err != nil {
		log.Panic("Cannot connect to DB", zap.Error(err))
	}
	defer db.Close()

	store := db.NewStore(pool)
	service := service.NewService(store)
	server, err := api.NewServer(*cfg, service)
	if err != nil {
		log.Panic("Failed to setup server", zap.Error(err))
	}

	if err := server.Start(":8080"); err != nil {
		log.Panic("Cannot start server", zap.Error(err))
	}
}
