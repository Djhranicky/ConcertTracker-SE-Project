package auth

import (
	"errors"
	"fmt"
	"net/http"
)

var cookieName string = "id"

func SetJWTCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    token,
		HttpOnly: true,
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

func VerifyJWTCookie(cookie *http.Cookie, inErr error) error {
	if cookie == nil {
		return inErr
	}
	tokenString := cookie.Value
	if tokenString == "" {
		return fmt.Errorf("missing authorization token")
	}
	err := VerifyToken(tokenString)
	if err != nil {
		return fmt.Errorf("invalid authorization token")
	}
	return nil
}
