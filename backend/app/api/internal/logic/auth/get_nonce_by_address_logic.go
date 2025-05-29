package authlogic

import (
	"context"
	"dex/app/api/internal/svc"
	"dex/app/api/internal/types"
	"math/rand"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetNonceByAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNonceByAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNonceByAddressLogic {
	return &GetNonceByAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNonceByAddressLogic) GetNonceByWalletAddress(req *types.GetNonceByAddressReq) (*types.GetNonceByAddressResp, error) {
	logx.Infof("GetNonceByWalletAddress req:%v", req)
	nonce, err := l.svcCtx.RedisClient.Get(l.ctx, req.WalletAddress).Result()
	if err != nil {
		if err.Error() == string(redis.Nil) {
			return nil, err
		}
	}
	if nonce == "" {
		id, err := l.svcCtx.SF.NextID()
		if err != nil {
			return nil, err
		}
		nonce, err := createNonce(req.WalletAddress, id)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.RedisClient.Set(context.Background(), req.WalletAddress, nonce, time.Hour*24).Err()
		if err != nil {
			return nil, err
		}
	}
	logx.Infof("GetNonceByWalletAddress address:%s, nonce:%s", req.WalletAddress, nonce)
	return &types.GetNonceByAddressResp{
		Nonce: nonce,
	}, nil
}

// 创建nonce
func createNonce(address string, id uint64) (string, error) {
	// 1. 将id转为36进制字符串（包含数字和小写字母）
	idStr := strconv.FormatUint(id, 36)
	// 2. 将address去掉0x前缀，取后8位
	addrPart := address
	if len(addrPart) > 2 && addrPart[:2] == "0x" {
		addrPart = addrPart[2:]
	}
	if len(addrPart) > 8 {
		addrPart = addrPart[len(addrPart)-8:]
	}
	// 3. 拼接后打乱顺序，取前10位
	raw := idStr + addrPart
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	runes := []rune(raw)
	r.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	// 取前10位，不足补0
	nonce := string(runes)
	if len(nonce) > 10 {
		nonce = nonce[:10]
	} else if len(nonce) < 10 {
		nonce = leftPad(nonce, "0", 10)
	}
	return nonce, nil
}

func leftPad(s, pad string, length int) string {
	for len(s) < length {
		s = pad + s
	}
	return s
}
