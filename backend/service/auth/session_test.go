package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
	"github.com/joho/godotenv"
)

func TestSessionMethods(t *testing.T) {
	utils.Init()
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("cannot load env file")
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	t.Run("should fail if request has no cookie", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		_, err = GetJWTCookie(req)
		if err != http.ErrNoCookie {
			t.Errorf("expected error code %v, got nothing", http.ErrNoCookie)
		}
	})

	t.Run("should pass if request has cookie", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    "",
			HttpOnly: true,
		})
		cookie, err := GetJWTCookie(req)
		if err != nil {
			t.Errorf("expected cookie in request, got %v", cookie)
		}
	})

	t.Run("verification should fail if no cookie present", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		err = VerifyJWTCookie(GetJWTCookie(req))
		if err != http.ErrNoCookie {
			t.Errorf("expected error code %v, got %v", http.ErrNoCookie, err)
		}
	})

	t.Run("verification should fail if no JWT token present", func(t *testing.T) {
		expectedErr := fmt.Errorf("missing authorization token")
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    "",
			HttpOnly: true,
		})
		err = VerifyJWTCookie(GetJWTCookie(req))
		if err.Error() != expectedErr.Error() {
			t.Errorf("expected error code %v, got %v", expectedErr, err)
		}
	})

	t.Run("verification should fail if JWT token is expired", func(t *testing.T) {
		expectedErr := fmt.Errorf("invalid authorization token")
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		token, _ := CreateJWT(secret, 1, -1)
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    token,
			HttpOnly: true,
		})
		err = VerifyJWTCookie(GetJWTCookie(req))
		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("expected error code %v, got %v", expectedErr, err)
		}
	})

	t.Run("verification should succeed if JWT token is valid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		token, _ := CreateJWT(secret, 1, 10)
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    token,
			HttpOnly: true,
		})
		err = VerifyJWTCookie(GetJWTCookie(req))
		if err != nil {
			t.Errorf("expected no error code, got %v", err)
		}
	})
}
