package auth

import (
	authlogic "dex/app/api/internal/logic/auth"
	"dex/app/api/internal/svc"
	"dex/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
)

// 获取钱包地址登录 nonce
func GetNonceByAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNonceByAddressReq
		err := httpx.Parse(r, &req)
		if err != nil {
			xhttp.JsonBaseResponse(w, err)
			return
		}
		l := authlogic.NewGetNonceByAddressLogic(r.Context(), svcCtx)
		resp, err := l.GetNonceByWalletAddress(&req)
		if err != nil {
			xhttp.JsonBaseResponse(w, err)
			return
		}

		xhttp.JsonBaseResponse(w, resp)
		return
	}
}
