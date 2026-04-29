package svc

import (
	"url-shortner/services/analytics-rpc/internal/config"
	"url-shortner/services/transform-rpc/transform"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	TransformerRpc transform.TransformerClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	zrpcClient := zrpc.MustNewClient(c.TransformerRpc)
	
	return &ServiceContext{
		Config: c,
		TransformerRpc: transform.NewTransformerClient(zrpcClient.Conn()),
	}
}
