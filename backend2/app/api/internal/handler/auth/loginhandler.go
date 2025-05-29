package auth

import (
	authlogic "dex/app/api/internal/logic/auth"
	"dex/app/api/internal/svc"
	"dex/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		err := httpx.ParseJsonBody(r, &req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		al := authlogic.NewAuthLogic(r.Context(), svcCtx)
		resp, err := al.Login(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
	}
}
