package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type PostgresConf struct {
	DSN string
}

type Config struct {
	zrpc.RpcServerConf
	Redis    redis.RedisConf
	Postgres PostgresConf
}
