package authRouter

import (
	"net/http"

	"github.com/pecet3/logger"
)

func (r router) handlePing(w http.ResponseWriter, req *http.Request) {
	u, err := r.app.Auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dto := u.ToDto(r.app.Data)

	err = dto.Send(w)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}
