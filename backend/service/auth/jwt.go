package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func CreateJWT(secret []byte, userID uint, seconds int) (string, error) {

	expiration := time.Second * time.Duration(seconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		strconv.Itoa(int(userID)),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		},
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, email string, store types.Store) error {
	user, err := store.GetUserByEmail(email)
	if err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return fmt.Errorf("unknown claims type")
	}

	claimsInt, err := strconv.Atoi(claims.UserID)

	if err != nil || uint(claimsInt) != user.ID {
		return fmt.Errorf("user not validated")
	}

	return nil
}
