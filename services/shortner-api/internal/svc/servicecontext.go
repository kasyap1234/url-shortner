// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"url-shortner/services/shortner-api/internal/config"
	"url-shortner/services/transform-rpc/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	TransformerClient transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		TransformerClient: transformer.NewTransformer(zrpc.MustNewClient(c.TransformerRpc)),
	}
}
