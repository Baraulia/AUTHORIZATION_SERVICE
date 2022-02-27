package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"time"
)

var logger = logging.GetLogger()

var Secret string

const AccessTokenTTL = time.Minute * 15
const RefreshTokenTTL = time.Hour * 24 * 30

func GenerateTokensByAuthUser(user *authProto.User) (*authProto.GeneratedTokens, error) {
	expired := time.Now().Add(AccessTokenTTL)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    user.UserId,
		"RoleId":    user.RoleId,
		"ExpiresAt": expired.Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(Secret))
	if err != nil {
		logger.Errorf("GenerateTokensByAuthUser:%s", err)
		return nil, fmt.Errorf("GenerateTokensByAuthUser:%w", err)
	}

	refreshExpired := time.Now().Add(RefreshTokenTTL)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": refreshExpired.Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(Secret))
	if err != nil {
		logger.Errorf("GenerateTokensByAuthUser:%s", err)
		return nil, fmt.Errorf("GenerateTokensByAuthUser:%w", err)
	}

	return &authProto.GeneratedTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
