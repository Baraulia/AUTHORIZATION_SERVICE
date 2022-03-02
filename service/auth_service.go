package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
	"strings"
	"time"
)

var Secret string

const AccessTokenTTL = time.Minute * 15
const RefreshTokenTTL = time.Hour * 24 * 30

type AuthService struct {
	logger logging.Logger
	repo   repository.Repository
}

func NewAuthService(repo repository.Repository, logger logging.Logger) *AuthService {
	return &AuthService{repo: repo, logger: logger}
}

func (a *AuthService) GenerateTokensByAuthUser(user *authProto.User) (*authProto.GeneratedTokens, error) {
	expired := time.Now().Add(AccessTokenTTL)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    user.UserId,
		"RoleId":    user.RoleId,
		"ExpiresAt": expired.Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(Secret))
	if err != nil {
		a.logger.Errorf("GenerateTokensByAuthUser:%s", err)
		return nil, fmt.Errorf("GenerateTokensByAuthUser:%w", err)
	}

	refreshExpired := time.Now().Add(RefreshTokenTTL)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    user.UserId,
		"ExpiresAt": refreshExpired.Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(Secret))
	if err != nil {
		a.logger.Errorf("GenerateTokensByAuthUser:%s", err)
		return nil, fmt.Errorf("GenerateTokensByAuthUser:%w", err)
	}

	return &authProto.GeneratedTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil

}

func (a *AuthService) ParseToken(token string) (*authProto.UserRole, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("ParseToken:%w", err)
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error while parsing access token")
	}
	if claims["ExpiresAt"].(int64) > time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	role, err := a.repo.RolePerm.GetRoleById(claims["RoleId"].(int))
	if err != nil {
		return nil, err
	}
	perms, err := a.repo.RolePerm.GetPermsByRoleId(claims["RoleId"].(int))
	if err != nil {
		return nil, err
	}
	var slicePerms []string
	for _, perm := range perms {
		slicePerms = append(slicePerms, perm.Name)
	}
	stringPerms := strings.Join(slicePerms[:], ",")
	return &authProto.UserRole{
		UserId:      claims["UserId"].(int32),
		Role:        role.Name,
		Permissions: stringPerms,
	}, nil
}

func (a *AuthService) RefreshTokens(refreshToken string) (*authProto.GeneratedTokens, error) {
	parseToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("RefreshTokens:%w", err)
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error while parsing access token")
	}
	if claims["ExpiresAt"].(int64) > time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	roleId, err := a.repo.RolePerm.GetRoleByUserId(claims["UserId"].(int))
	if err != nil {
		return nil, err
	}
	return a.GenerateTokensByAuthUser(&authProto.User{UserId: claims["UserId"].(int32), RoleId: int32(roleId)})
}

func (a *AuthService) CheckRights(token string, requiredRole string) (bool, error) {
	userRole, err := a.ParseToken(token)
	if err != nil {
		return false, err
	}
	if userRole.Role != requiredRole {
		return false, errors.New("no required rights")
	}
	return true, nil
}
