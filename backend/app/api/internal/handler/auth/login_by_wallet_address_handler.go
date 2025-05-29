package auth

import (
	authlogic "dex/app/api/internal/logic/auth"
	"dex/app/api/internal/svc"
	"dex/app/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func LoginByWalletAddress(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WalletAddressLoginReq
		err := httpx.Parse(r, &req)
		if err != nil {
			xhttp.JsonBaseResponse(w, err)
			return
		}
		l := authlogic.NewWalletAddressLoginLogic(r.Context(), svcCtx)
		resp, err := l.WalletAddressLogin(&req)
		if err != nil {
			xhttp.JsonBaseResponse(w, err)
			return
		}

		xhttp.JsonBaseResponse(w, resp)
		return
	}
}
