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
		Name:     cookieName,
		Value:    token,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
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

func VerifyJWTCookie(cookie *http.Cookie, userID uint) error {
	tokenString := cookie.Value
	if tokenString == "" {
		return fmt.Errorf("missing authorization token")
	}
	err := VerifyToken(tokenString, userID)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUser(r *http.Request, store types.Store) error {
	cookie, err := GetJWTCookie(r)
	if err != nil {
		return err
	}

	username := r.Header.Get("username")
	if username == "" {
		return fmt.Errorf("no username provided")
	}

	user, err := store.GetUserByUsername(username)
	if err != nil {
		return err
	}

	return VerifyJWTCookie(cookie, user.ID)
}
