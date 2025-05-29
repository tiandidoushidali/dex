package authlogic

import (
	"dex/app/api/internal/svc"
	"dex/app/api/internal/types"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"golang.org/x/net/context"

	"github.com/ethereum/go-ethereum/common"
)

type LoginByWalletAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWalletAddressLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByWalletAddressLogic {
	return &LoginByWalletAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByWalletAddressLogic) WalletAddressLogin(req *types.WalletAddressLoginReq) (resp *types.JwtAuthResp, err error) {
	// 1.获取nonce
	nonce, err := l.svcCtx.RedisClient.Get(l.ctx, "wallet_address_login_"+req.WalletAddress).Result()
	if err != nil {
		if err.Error() == string(redis.Nil) {
			return nil, errors.New(-1, "Nonce not exist")
		}
		logx.WithContext(l.ctx).Errorf("get nonce from redis err: %+v", err)
		return nil, errors.New(-1, "Nonce not exist")
	}
	// 2.校验签名
	err = validateEthWallet(req.WalletAddress, nonce, req.Signature)
	if err != nil {
		return nil, err
	}
	// 查询用户
	return nil, err
}

func validateEthWallet(address, nonce, signature string) error {
	logx.Infof("validateEthWallet address:%s nonce:%s signature:%s", address, nonce, signature)
	addrKey := common.HexToAddress(address)
	sig := hexutil.MustDecode(signature)
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}
	msg := fmt.Sprint("\x19Etherum Signed Message:\n%d%s", len(nonce), nonce)
	msg256 := crypto.Keccak256([]byte(msg))
	pubKey, err := crypto.SigToPub(msg256, sig)
	if err != nil {
		return err
	}
	recoverAddr := crypto.PubkeyToAddress(*pubKey)
	if recoverAddr != addrKey {
		return errors.New(-1, "Address mismatch")
	}

	return nil
}
