package authRouter

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/pecet3/las-test-pdf/data/dtos"
	"github.com/pecet3/logger"
)

func (r router) handleLogin(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.Login{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()
	u, err := r.app.Data.GetUserByEmail(ctx, sql.NullString{
		String: dto.Email,
		Valid:  true,
	})
	if u.ID == 0 || err != nil {
		logger.Error(err)
		http.Error(w, "User with provided email doesn't exists", http.StatusBadRequest)
		return
	}
	s, code := r.app.Auth.New.LoginSession(u.ID, u.Uuid)
	if err := r.app.Auth.Sessions["login"].Add(s); err != nil {
		logger.WarnC(err, " User ID: ", u.ID)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.app.Auth.Mailers["login"].Send(ctx, string(u.Email.String), code, u.Name)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.InfoC(fmt.Sprintf(`<Login> User with email: %s Access Code: %s`, u.Email.String, s.ActivateCode))
}
