// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"

	"url-shortner/services/shortner-api/internal/svc"
	"url-shortner/services/shortner-api/internal/types"
	"url-shortner/services/transform-rpc/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectLogic {
	return &RedirectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectLogic) Redirect(req *types.RedirectRequest) (resp *types.RedirectResponse, err error) {
	rpcResp, err := l.svcCtx.TransformerClient.Expand(l.ctx, &transformer.ExpandRequest{
		ShortUrl: req.ShortURL,
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.Url == "" {
		return nil, errors.New("url not found")
	}

	return &types.RedirectResponse{URL: rpcResp.Url}, nil
}
