package handler

import (
	"chat/internal/logic"
	"chat/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		resp, err := logic.UserLogin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
