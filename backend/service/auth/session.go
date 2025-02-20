package auth

import "net/http"

func setJWTCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "id",
		Value:    token,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}
