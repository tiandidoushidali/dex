package authlogic

import (
	"context"
	"dex/app/api/internal/svc"
	"dex/app/api/internal/types"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(8)).Unix()
	claims["addr"] = req.Username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))

	resp := &types.LoginResp{Token: token.Signature}
	return resp, nil
}
