package authRouter

import (
	"net/http"

	"github.com/pecet3/logger"
)

func (r router) handleLogout(w http.ResponseWriter, req *http.Request) {
	s, err := r.app.Auth.GetContextSession(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if err := r.app.Auth.UpdateSessionIsExpired(s.Token); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	r.app.Auth.ClearCookies(w)
}
