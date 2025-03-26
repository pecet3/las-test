package authRouter

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/pecet3/las-test-pdf/data/dtos"
	"github.com/pecet3/logger"
)

func (r router) handleLoginExchange(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.Exchange{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	u, err := r.app.Data.GetUserByEmail(req.Context(), sql.NullString{String: dto.Email, Valid: true})
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	_, code, err := r.app.Auth.Sessions["login"].Get(u.Uuid)
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if code != dto.Code {
		logger.WarnC(fmt.Sprintf("User with Email: %s provided wrong token during login", dto.Email))
		http.Error(w, "Invalid Code", http.StatusBadRequest)
		return
	}

	ns, token, err := r.app.Auth.New.AuthSession(u.ID, u.Email.String, u.Name)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	s, err := r.app.Auth.AddAuthSession(token, ns)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.InfoC("New Auth Session: ", s)

	r.app.Auth.SetCookies(w, &s)

	r.app.Auth.Sessions["login"].Remove(u.Uuid)
}
