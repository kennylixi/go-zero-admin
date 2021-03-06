package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendSubjectDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectDeleteLogic {
	return &HomeRecommendSubjectDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendSubjectDeleteLogic) HomeRecommendSubjectDelete(in *sms.HomeRecommendSubjectDeleteReq) (*sms.HomeRecommendSubjectDeleteResp, error) {
	err := l.svcCtx.SmsHomeRecommendSubjectModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendSubjectDeleteResp{}, nil
}
