package auth

import (
	"net/http"
	"time"

	"github.com/pecet3/las-test-pdf/data"
)

func (a *Auth) SetCookies(w http.ResponseWriter, s *data.Session) {
	cookie := http.Cookie{
		Name:     "auth",
		Value:    s.Token,
		Expires:  time.Now().Add(time.Hour * 192),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	cookieRefresh := http.Cookie{
		Name:     "refresh",
		Value:    s.RefreshToken,
		Expires:  time.Now().Add(time.Hour * 192 * 2),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookieRefresh)

}

func (a *Auth) ClearCookies(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Ustawienie czasu wygaśnięcia w przeszłości
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
}
