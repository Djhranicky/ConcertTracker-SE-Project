package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
)

var cookieName string = "id"

func SetJWTCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:  cookieName,
		Value: token,
	}

	http.SetCookie(w, &cookie)
}

func GetJWTCookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		var returnErr error
		switch {
		case errors.Is(err, http.ErrNoCookie):
			returnErr = err
		default:
			returnErr = fmt.Errorf("server error getting JWT Cookie")
		}
		return nil, returnErr
	}
	return cookie, nil
}

func VerifyJWTCookie(cookie *http.Cookie, inErr error, email string, store types.Store) error {
	if cookie == nil {
		return inErr
	}
	tokenString := cookie.Value
	if tokenString == "" {
		return fmt.Errorf("missing authorization token")
	}
	err := VerifyToken(tokenString, email, store)
	if err != nil {
		return err
	}
	return nil
}
