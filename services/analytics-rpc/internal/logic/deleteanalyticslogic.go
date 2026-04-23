package logic

import (
	"context"

	"url-shortner/services/analytics-rpc/analytics"
	"url-shortner/services/analytics-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAnalyticsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAnalyticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAnalyticsLogic {
	return &DeleteAnalyticsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAnalyticsLogic) DeleteAnalytics(in *analytics.AnalyticsRequest) (*analytics.AnalyticsResponse, error) {
	// todo: add your logic here and delete this line

	return &analytics.AnalyticsResponse{}, nil
}
