package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("jwt_secret")

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	rtExpireTime := nowTime.Add(10 * 24 * time.Hour)
	claims := Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Ekioo_Mall",
		},
	}
	// 加密并获得完整的编码后的字符串token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "mall",
	}).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// ParseRefreshToken 验证用户token
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return
	}

	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return
	}

	if accessClaim.ExpiresAt > time.Now().Unix() {
		// 如果 access_token 没过期,每一次请求都刷新 refresh_token 和 access_token
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}

	if refreshClaim.ExpiresAt > time.Now().Unix() {
		// 如果 access_token 过期了,但是 refresh_token 没过期, 刷新 refresh_token 和 access_token
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}

	// 如果两者都过期了,重新登陆
	return "", "", errors.New("身份过期，重新登陆")
}
