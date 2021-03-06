package auth_util

import (
	"github.com/huangxuantao/utils/time_util"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func getClaims() CustomClaims {
	claims := CustomClaims{
		Username: "admin",
		NowTime:  time_util.GetNowTime(),
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Truncate(1 * time.Hour).Unix()),
			ExpiresAt: int64(time.Now().Add(1 * time.Hour).Unix()),
			Issuer:    "veryon-worker",
		},
	}
	return claims
}

func TestAuthJWT_CreateToken(t *testing.T) {
	j := NewJWT()
	claims := getClaims()
	if token, err := j.CreateToken(claims); err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}

}
