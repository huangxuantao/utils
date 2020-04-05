package auth_util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("token is illegal")
	signKey          = "5a709f0d2a5e9f2877e318a82527fb68"

	ExpireTime = 7 * 24 * time.Hour
)

type Info struct {
	ID       string   `json:"id"`       // ID标识
	Username string   `json:"username"` // 客户端名称
	IP       string   `json:"ip"`       // 客户端申请token时的IP地址
	Sub      []string `json:"sub"`      // 客户端允许访问的资源 先不判断允许访问的资源
	NowTime  string   `json:"now_time"`
}

type CustomClaims struct {
	Info
	jwt.StandardClaims
}

type authJWT struct {
	SignKey []byte
}

func NewJWT() *authJWT {
	return &authJWT{
		SignKey: []byte(signKey),
	}
}

func (j *authJWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

func (j *authJWT) ParseToken(tokenString string) (*CustomClaims, error) {
	claims := new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return j.SignKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *authJWT) RefreshToken(tokenString string) (string, error) {
	claims := new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return j.SignKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		claims.StandardClaims.ExpiresAt = time.Now().Add(ExpireTime).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func GetUsername(tokenString string) (string, error) {
	authJWT := NewJWT()
	claims, err := authJWT.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Username, nil
}

func GetInfo(tokenString string) (*Info, error) {
	authJWT := NewJWT()
	claims, err := authJWT.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return &claims.Info, nil
}
