package smsprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSmsProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSmsProviderLogic {
	return &UpdateSmsProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSmsProviderLogic) UpdateSmsProvider(req *types.SmsProviderInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.UpdateSmsProvider(l.ctx,
		&mcms.SmsProviderInfo{
			Id:        req.Id,
			Name:      req.Name,
			SecretId:  req.SecretId,
			SecretKey: req.SecretKey,
			Region:    req.Region,
			IsDefault: req.IsDefault,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
