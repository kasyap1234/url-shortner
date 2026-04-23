package logic

import (
	"context"
	"errors"

	"url-shortner/services/transform-rpc/internal/svc"
	"url-shortner/services/transform-rpc/transform"

	"github.com/jackc/pgx/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandRequest) (*transform.ExpandResponse, error) {
	url, err := l.svcCtx.Redis.Get(in.ShortUrl)
	if err == nil && url != "" {
		return &transform.ExpandResponse{Url: url}, nil
	}

	urlMap, err := l.svcCtx.Store.GetURLByShortCode(l.ctx, in.ShortUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("url not found")
		}

		return nil, err
	}

	if err := l.svcCtx.Redis.Setex(in.ShortUrl, urlMap.Url, 604800); err != nil {
		l.Errorf("failed to warm cache for short url %s: %v", in.ShortUrl, err)
	}

	return &transform.ExpandResponse{Url: urlMap.Url}, nil
}
