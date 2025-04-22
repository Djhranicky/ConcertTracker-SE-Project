package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
)

var cookieName string = "id"

func SetJWTCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   3600 * 24,
	}

	http.SetCookie(w, &cookie)
}

func GetJWTCookie(r *http.Request) (*http.Cookie, error) {
	for _, cookie := range r.Cookies() {
		log.Printf("%v=%v\n", cookie.Name, cookie.Value)
	}
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		var returnErr error
		switch {
		case errors.Is(err, http.ErrNoCookie):
			log.Println("No cookie :(")
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
