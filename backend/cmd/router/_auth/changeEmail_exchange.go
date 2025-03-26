package authRouter

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/las-test-pdf/data/dtos"
	"github.com/pecet3/logger"
)

func (r router) handleChangeEmailExchange(w http.ResponseWriter, req *http.Request) {
	u, err := r.app.Auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dto := &dtos.Exchange{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	ne, code, err := r.app.Auth.Sessions["change_email"].Get(u.Uuid)
	newEmail := ne.(string)
	r.app.Data.UpdateUserEmail(req.Context(), data.UpdateUserEmailParams{
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		ID: u.ID,
	})
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if code != dto.Code {
		logger.WarnC(fmt.Sprintf("User with Email: %s provided wrong token during change", dto.Email))
		http.Error(w, "Invalid Code", http.StatusBadRequest)
		return
	}
	s, err := r.app.Auth.GetContextSession(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "No session found", http.StatusBadRequest)
		return
	}
	r.app.Auth.Sessions["change_email"].Remove(u.Uuid)

	// Logout
	r.app.Auth.UpdateSessionIsExpired(s.Token)
	r.app.Auth.ClearCookies(w)
}
