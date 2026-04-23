package logic

import (
	"context"

	"url-shortner/services/analytics-rpc/analytics"
	"url-shortner/services/analytics-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnalyticsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnalyticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnalyticsLogic {
	return &GetAnalyticsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnalyticsLogic) GetAnalytics(in *analytics.AnalyticsRequest) (*analytics.AnalyticsResponse, error) {
	// todo: add your logic here and delete this line

	return &analytics.AnalyticsResponse{}, nil
}
