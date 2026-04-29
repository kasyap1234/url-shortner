package svc

import (
	"context"
	"time"

	"url-shortner/services/transform-rpc/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config    config.Config
	Redis     *redis.Redis 
	Generator *sonyflake.Sonyflake
}

func NewServiceContext(c config.Config) *ServiceContext {
	
	
	return &ServiceContext{
		Config:    c,
		Redis:     redis.MustNewRedis(c.Redis),
		Generator: sonyflake.NewSonyflake(sonyflake.Settings{}),
	}
}

func mustNewPostgres(dsn string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		panic(err)
	}

	return db
}
