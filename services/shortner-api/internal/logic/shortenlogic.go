// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"url-shortner/services/shortner-api/internal/svc"
	"url-shortner/services/shortner-api/internal/types"
	"url-shortner/services/transform-rpc/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenRequest) (resp *types.ShortenResponse, err error) {
	rpcResp, err := l.svcCtx.TransformerClient.Shorten(l.ctx, &transformer.ShortenRequest{
		Url: req.URL,
	})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResponse{ShortURL: rpcResp.ShortUrl}, nil
}
