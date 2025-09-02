package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var dbPool *pgxpool.Pool

func Connect(ctx context.Context, databaseURL string, log *zap.Logger) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Error("cannot create db pool", zap.Error(err))
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Error("cannot ping db", zap.Error(err))
		return nil, err
	}

	log.Info("connected to db successfully")
	return pool, nil
}

// Close cleans up the DB connection pool
func Close() {
	if dbPool != nil {
		dbPool.Close()
	}
}
