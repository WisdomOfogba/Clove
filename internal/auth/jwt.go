// Have chosen to go with Redis for session management. (for now)
package auth

import (
	"fmt"
	"time"

	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/global"

	"github.com/golang-jwt/jwt/v5"
)

type JWTField struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// expirationTime is the time in seconds until the token expires.
func GenerateBackendAccessToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTField{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    constants.APPNAME,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.UserAccessTkDur)),
		},
	})

	tokenString, err := token.SignedString(global.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateBackendAccessToken(tokenString string, secretKey []byte) (*JWTField, error) {
	logger := global.Logger
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate backend jwt token")
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jwtField := JWTField{}
		claimsToJWTField(claims, &jwtField)
		// userId := jwtField.UserID
		return &jwtField, nil
	}

	return nil, jwt.ErrTokenSignatureInvalid
}

func claimsToJWTField(claims jwt.MapClaims, jwtField *JWTField) {
	for k, v := range claims {
		switch k {
		case "user_id":
			userId, _ := v.(int64)
			jwtField.UserID = userId
		case "iss":
			issuer, _ := v.(string)
			jwtField.Issuer = issuer
		case "exp":
			// expiryStr, _ := strconv.ParseInt(v.(int64), 10, 64)
			expiryInt, _ := v.(int64)
			expiry := time.Unix(expiryInt, 0)
			jwtField.ExpiresAt = jwt.NewNumericDate(expiry)
		case "iat":
			// expiryStr, _ := strconv.ParseInt(v.(int64), 10, 64)
			issuedAtInt, _ := v.(int64)
			issuedAt := time.Unix(issuedAtInt, 0)
			jwtField.ExpiresAt = jwt.NewNumericDate(issuedAt)
		}
	}
}

// func validateJWTMeta(jwtStruct *JWTField) bool {
// 	if jwtStruct.Issuer != constants.APPNAME {
// 		return false
// 	}

// 	if time.Now().After(jwtStruct.ExpiresAt.Time) {
// 		return false
// 	}

// 	return true
// }
