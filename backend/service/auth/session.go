package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
)

func setJWTCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "id",
		Value:    token,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func getJWTCookie(w http.ResponseWriter, r http.Request) *http.Cookie {
	cookie, err := r.Cookie("id")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			utils.WriteError(w, http.StatusBadRequest, err)
		default:
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("server error getting JWT Cookie"))
		}
		return nil
	}
	return cookie
}
