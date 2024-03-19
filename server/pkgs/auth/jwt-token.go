package tokens

import (
	"fmt"
	"projects/chatterbox/server/pkgs/dao"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Role      string `json:"role"`
	UserId    int64  `json:"user_id"`
	SessionId string `json:"session_id"`
	jwt.StandardClaims
}

const (
	JwtSecretKey = "my_secret_key"
)

var (
	JWtExpiry = time.Now().Add(24 * time.Hour) // 1 day
	JwtKey    = []byte(JwtSecretKey)
)

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func SignToken(user dao.User) (string, string, int, error) {
	sessionId := uuid.New().String()
	claims := &Claims{
		Role:      user.Role,
		UserId:    user.UserId,
		SessionId: sessionId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "www.chatterbox.in",
			Audience:  "chatterboxes",
			Subject:   fmt.Sprintf("chatterbox#%v#%v", user.Role, user.UserId),
			ExpiresAt: JWtExpiry.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	return sessionId, tokenString, int(JWtExpiry.Unix()), err
}
