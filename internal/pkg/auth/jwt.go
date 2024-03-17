package auth

import (
	"fmt"
	"github.com/spf13/viper"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Uid      uint   `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// generate tokens used for auth
func GenerateToken(userID int64, userName string) string {

	claims := Claims{
		uint(userID),
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt64("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "jwt", // 签发人
		},
	}

	//claims.ExpiresAt = time.Now().Add(time.Duration(configs.Conf.AuthConfig.JwtExpire) * time.Hour).Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(viper.GetString("auth.jwt_secret")))

	if err != nil {
		panic(err)
	}
	return token
}

// verify token
func JwtVerify(tokenStr string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("auth.jwt_secret")), nil
	})

	if !token.Valid || err != nil {
		return nil, fmt.Errorf("token invalid")
	}
	claims, ok := token.Claims.(*Claims)

	if float64(claims.ExpiresAt) < float64(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	if !ok {
		return nil, err
	}
	return claims, err

}
