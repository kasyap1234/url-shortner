package logic

import (
	"context"

	"url-shortner/services/transform-rpc/internal/store"
	"url-shortner/services/transform-rpc/internal/svc"
	"url-shortner/services/transform-rpc/transform"

	"github.com/jcoene/go-base62"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenRequest) (*transform.ShortenResponse, error) {
	id, err := l.svcCtx.Generator.NextID()
	if err != nil {
		return nil, err
	}
	shortCode := base62.EncodeUint64(id)

	err = l.svcCtx.Store.CreateURLMap(l.ctx, store.CreateURLMapParams{
		ShortCode: shortCode,
		Url:       in.Url,
	})
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.Redis.Setex(shortCode, in.Url, 604800); err != nil {
		l.Errorf("failed to cache short url %s: %v", shortCode, err)
	}

	return &transform.ShortenResponse{ShortUrl: shortCode}, nil
}
