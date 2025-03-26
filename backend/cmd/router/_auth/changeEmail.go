package authRouter

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pecet3/las-test-pdf/data/dtos"
	"github.com/pecet3/logger"
)

func (r router) handleChangeEmail(w http.ResponseWriter, req *http.Request) {
	u, err := r.app.Auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dto := &dtos.EmailChange{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()
	if dto.Email == u.Email.String {
		http.Error(w, "Provided now new email address", http.StatusBadRequest)
		return
	}
	s, code := r.app.Auth.New.ChangeEmailSession(u.ID, u.Uuid, dto.Email)
	if err := r.app.Auth.Sessions["change_email"].Add(s); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.app.Auth.Mailers["change_email"].Send(ctx, string(u.Email.String), code, u.Name)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.InfoC(fmt.Sprintf(`<ChangeEmail> User with email: %s Access Code: %s`, u.Email.String, s.ActivateCode))
}
