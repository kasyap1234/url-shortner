package logic

import (
	"context"

	"url-shortner/services/analytics-rpc/analytics"
	"url-shortner/services/analytics-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAnalyticsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAnalyticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAnalyticsLogic {
	return &CreateAnalyticsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAnalyticsLogic) CreateAnalytics(in *analytics.AnalyticsRequest) (*analytics.AnalyticsResponse, error) {
	// todo: add your logic here and delete this line

	return &analytics.AnalyticsResponse{}, nil
}
