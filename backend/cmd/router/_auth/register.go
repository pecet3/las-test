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

func (r router) handleRegister(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.Register{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()

	existingUser, err := r.app.Data.GetUserByEmail(ctx, sql.NullString{String: dto.Email, Valid: true})
	if existingUser.ID != 0 || err == nil {
		if !existingUser.IsDraft {
			logger.WarnC("Attempt to create account with existing email: ", existingUser.Email)
			http.Error(w, "User with provided email already exists", http.StatusBadRequest)
			return
		}
		s, code := r.app.Auth.New.RegisterSession(existingUser.ID, existingUser.Uuid)
		if err := r.app.Auth.Sessions["register"].Add(s); err != nil {
			logger.WarnC(err, " User ID: ", existingUser.ID)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := r.app.Auth.Mailers["register"].Send(ctx, dto.Email, code, dto.Name); err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		logger.InfoC(fmt.Sprintf(`<Register Resend> User with email: %s . Access Code: %s`,
			existingUser.Email.String, s.ActivateCode))

		return
	}
	u, err := r.app.Auth.AddUserTablesDb(dto.Name, dto.Email)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	s, code := r.app.Auth.New.RegisterSession(u.ID, u.Uuid)
	if err := r.app.Auth.Sessions["register"].Add(s); err != nil {
		logger.WarnC(err, " User ID: ", u.ID)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := r.app.Auth.Mailers["register"].Send(ctx, dto.Email, code, dto.Name); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.InfoC(fmt.Sprintf(`<Register> User with email: %s . Access Code: %s`, u.Email.String, s.ActivateCode))

}
