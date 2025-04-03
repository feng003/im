package ctxdata

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessTokenExpiredTime  = 60 * 60 // 5 hours
	RefreshTokenExpiredTime = 7 * 24 * 3600
	AccessTokenType         = "x-access"  // 5 minutes
	RefreshTokenType        = "x-refresh" // 30 days
	AuthSecret              = "asdasdasdAuthSecretasdsdasd"
)

func GetFrontAccessToken(secret string, expire int64, tokenType string, uid, msId, storeId int64) (string, int64) {

	now := time.Now().Unix()
	tokenContent := jwt.MapClaims{
		"iat":           now,
		"exp":           now + expire,
		"type":          tokenType,
		FrontJwtUserId:  uid,
		FrontJwtMsId:    msId,
		FrontJwtStoreId: storeId,
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", 0
	}

	return token, now + expire
}

func GetCmsAccessToken(secret string, expire int64, tokenType string, uid, msId, storeId int64) (string, int64) {

	now := time.Now().Unix()
	tokenContent := jwt.MapClaims{
		"iat":         now,
		"exp":         now + expire,
		"type":        tokenType,
		CmsJwtUserId:  uid,
		CmsJwtMsId:    msId,
		CmsJwtStoreId: storeId,
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", 0
	}

	return token, now + expire
}

//func GenerateAccessToken(payload map[string]interface{}) string {
//	payload["type"] = AccessTokenType
//	tokenContent := jwt.MapClaims{
//		"payload": payload,
//		"exp":     time.Now().Add(time.Second * AccessTokenExpiredTime).Unix(),
//	}
//	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
//	token, err := jwtToken.SignedString([]byte(AuthSecret))
//	if err != nil {
//		logx.Errorf("Failed to generate access token: %+v \n", err)
//		return ""
//	}
//
//	return token
//}
//
//func GenerateRefreshToken(payload map[string]interface{}) string {
//
//	payload["type"] = RefreshTokenType
//	tokenContent := jwt.MapClaims{
//		"payload": payload,
//		"exp":     time.Now().Add(time.Second * RefreshTokenExpiredTime).Unix(),
//	}
//	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
//	token, err := jwtToken.SignedString([]byte(AuthSecret))
//	if err != nil {
//		logx.Errorf("Failed to generate refresh token: %+v \n", err)
//		return ""
//	}
//
//	return token
//}

//func RefreshAccessToken(refreshTokenString string) (string, error) {
//	// 验证刷新 token
//	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(AuthSecret), nil
//	})
//	if err != nil {
//		return "", err
//	}
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok || !token.Valid {
//		return "", jwt.ErrInvalidKey
//	}
//
//	// 检查刷新 token 的类型
//	payload, ok := claims["payload"].(map[string]interface{})
//	if !ok || payload["type"] != RefreshTokenType {
//		return "", fmt.Errorf("invalid refresh token type")
//	}
//
//	// 生成新的访问 token 使用刷新 token 的 payload
//	newAccessToken := GenerateAccessToken(payload)
//	if newAccessToken == "" {
//		return "", fmt.Errorf("failed to generate new access token")
//	}
//
//	return newAccessToken, nil
//}
//
//func ValidateToken(jwtToken string) (map[string]interface{}, error) {
//
//	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
//	tokenData := jwt.MapClaims{}
//	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
//		return []byte(AuthSecret), nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	if !token.Valid {
//		return nil, jwt.ErrInvalidKey
//	}
//
//	var data map[string]interface{}
//	_ = copier.Copy(&data, tokenData["payload"])
//
//	return data, nil
//}
