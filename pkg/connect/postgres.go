package connect

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Postgres(ctx context.Context, dbURL string, log *zap.Logger) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Panic("cannot parse db config", zap.Error(err))
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Panic("cannot create db pool", zap.Error(err))
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Panic("cannot ping db", zap.Error(err))
		return nil, err
	}

	log.Info("connected to db successfully")
	return pool, nil
}
